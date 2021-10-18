package reporter

import (
	"compress/gzip"
	"context"
	"crypto/ecdsa"
	"crypto/x509"
	_ "embed"
	"encoding/csv"
	"encoding/pem"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/deepmap/oapi-codegen/pkg/securityprovider"
	"github.com/gocarina/gocsv"
	"github.com/golang-jwt/jwt"
	"github.com/jinzhu/now"
	"github.com/morishin/appstore-connect-sales-reporter/openapi"
)

func getSalesReports(accessInfo *AppStoreConnectAPIAccessInfo) SalesReports {
	dayBeforeYesterday := now.BeginningOfDay().AddDate(0, 0, -2)
	lastWeek := now.With(dayBeforeYesterday).BeginningOfWeek()
	lastMonth := now.With(dayBeforeYesterday.AddDate(0, -1, 0)).BeginningOfMonth()
	lastYear := now.With(dayBeforeYesterday.AddDate(-1, 0, 0)).BeginningOfYear()

	var dayBeforeYesterdayReport *SalesReport
	var lastWeekReport *SalesReport
	var lastMonthReport *SalesReport
	var lastYearReport *SalesReport

	jwtStr := generateJWT(accessInfo)
	wg := sync.WaitGroup{}
	wg.Add(4)
	go func() {
		defer wg.Done()
		dayBeforeYesterdayReport = getSalesReport(accessInfo.BaseUrl, jwtStr, dayBeforeYesterday.Format("2006-01-02"), "DAILY")
	}()
	go func() {
		defer wg.Done()
		lastWeekReport = getSalesReport(accessInfo.BaseUrl, jwtStr, lastWeek.Format("2006-01-02"), "WEEKLY")
	}()
	go func() {
		defer wg.Done()
		lastMonthReport = getSalesReport(accessInfo.BaseUrl, jwtStr, lastMonth.Format("2006-01"), "MONTHLY")
	}()
	go func() {
		defer wg.Done()
		lastYearReport = getSalesReport(accessInfo.BaseUrl, jwtStr, lastYear.Format("2006"), "YEARLY")
	}()
	wg.Wait()

	return SalesReports{
		DayBeforeYesterday: *dayBeforeYesterdayReport,
		LastWeek:           *lastWeekReport,
		LastMonth:          *lastMonthReport,
		LastYear:           *lastYearReport,
	}
}

func salesReportsToProceeds(salesReports *SalesReports, currency string) Proceeds {
	calcProceeds := func(salesReport *SalesReport, currency string) int {
		result := 0
		for _, row := range *salesReport {
			if row.CurrencyOfProceeds == currency {
				proceeds, err1 := strconv.ParseFloat(row.DeveloperProceeds, 32)
				if err1 != nil {
					panic(err1)
				}
				units, err2 := strconv.ParseFloat(row.Units, 32)
				if err2 != nil {
					panic(err2)
				}
				result += int(proceeds * units)
			}
		}
		return result
	}

	return Proceeds{
		DayBeforeYesterday: calcProceeds(&salesReports.DayBeforeYesterday, currency),
		LastWeek:           calcProceeds(&salesReports.LastWeek, currency),
		LastMonth:          calcProceeds(&salesReports.LastMonth, currency),
		LastYear:           calcProceeds(&salesReports.LastYear, currency),
	}
}

func generateJWT(accessInfo *AppStoreConnectAPIAccessInfo) string {
	expireTime := time.Now().Add(time.Minute * 10).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodES256, jwt.MapClaims{
		"iss": accessInfo.IssuerId,
		"exp": expireTime,
		"aud": "appstoreconnect-v1",
	})
	token.Header["kid"] = accessInfo.KeyID

	key, err := readPrivateKey()
	if err != nil {
		panic(err)
	}
	jwtStr, err := token.SignedString(key)
	if err != nil {
		panic(err)
	}
	return jwtStr
}

func getSalesReport(baseUrl string, jwtStr string, reportDate string, frequency string) *SalesReport {
	bearerTokenProvider, bearerTokenProviderErr := securityprovider.NewSecurityProviderBearerToken(jwtStr)
	if bearerTokenProviderErr != nil {
		panic(bearerTokenProviderErr)
	}
	client, clientErr := openapi.NewClient(baseUrl, openapi.WithRequestEditorFn(bearerTokenProvider.Intercept))
	if clientErr != nil {
		panic(clientErr)
	}

	setAcceptHeader := func(ctx context.Context, req *http.Request) error {
		req.Header.Set("Accept", "application/json, application/a-gzip")
		return nil
	}
	res, resErr := client.SalesReportsGetCollection(context.Background(), &openapi.SalesReportsGetCollectionParams{
		FilterFrequency:     []openapi.SalesReportsGetCollectionParamsFilterFrequency{openapi.SalesReportsGetCollectionParamsFilterFrequency(frequency)},
		FilterReportDate:    &[]string{reportDate},
		FilterReportSubType: []openapi.SalesReportsGetCollectionParamsFilterReportSubType{"SUMMARY"},
		FilterReportType:    []openapi.SalesReportsGetCollectionParamsFilterReportType{"SALES"},
		FilterVendorNumber:  []string{"85696015"},
	}, setAcceptHeader)
	if resErr != nil {
		panic(resErr)
	}

	result := unmarshalSalesReport(res.Body)
	return &result
}

//go:embed AuthKey.p8
var authKeyP8 []byte

func readPrivateKey() (*ecdsa.PrivateKey, error) {
	var err error
	var parsedKey interface{}
	var key *ecdsa.PrivateKey
	var ok bool
	block, _ := pem.Decode(authKeyP8)
	if parsedKey, err = x509.ParsePKCS8PrivateKey(block.Bytes); err != nil {
		return nil, err
	}
	if key, ok = parsedKey.(*ecdsa.PrivateKey); !ok {
		return nil, fmt.Errorf("not a EC private key file")
	}
	return key, nil
}

func unmarshalSalesReport(gzipFile io.ReadCloser) []*SalesReportRow {
	tempFile, err := os.CreateTemp("", "")
	if err != nil {
		panic(err)
	}
	ungzip, err := gzip.NewReader(gzipFile)
	if err != nil {
		panic(err)
	}
	io.Copy(tempFile, ungzip)
	salesReport := []*SalesReportRow{}
	gocsv.SetCSVReader(func(in io.Reader) gocsv.CSVReader {
		r := csv.NewReader(in)
		r.Comma = '\t'
		return r
	})
	_, err = tempFile.Seek(0, io.SeekStart)
	if err != nil {
		panic(err)
	}
	if err := gocsv.Unmarshal(tempFile, &salesReport); err != nil {
		panic(err)
	}
	return salesReport
}
