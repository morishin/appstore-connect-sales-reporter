package main

type Proceeds struct {
	DayBeforeYesterday int
	LastWeek           int
	LastMonth          int
	LastYear           int
}
type SalesReports struct {
	DayBeforeYesterday SalesReport
	LastWeek           SalesReport
	LastMonth          SalesReport
	LastYear           SalesReport
}
type SalesReport = []*SalesReportRow
type SalesReportRow struct {
	Provider              string `csv:"Provider"`
	ProviderCountry       string `csv:"Provider Country"`
	Sku                   string `csv:"SKU"`
	Developer             string `csv:"Developer"`
	Title                 string `csv:"Title"`
	Version               string `csv:"Version"`
	ProductTypeIdentifier string `csv:"Product Type Identifier"`
	Units                 string `csv:"Units"`
	DeveloperProceeds     string `csv:"Developer Proceeds"`
	BeginDate             string `csv:"Begin Date"`
	EndDate               string `csv:"End Date"`
	CustomerCurrency      string `csv:"Customer Currency"`
	CountryCode           string `csv:"Country Code"`
	CurrencyOfProceeds    string `csv:"Currency of Proceeds"`
	AppleIdentifier       string `csv:"Apple Identifier"`
	CustomerPrice         string `csv:"Customer Price"`
	PromoCode             string `csv:"Promo Code"`
	ParentIdentifier      string `csv:"Parent Identifier"`
	Subscription          string `csv:"Subscription"`
	Period                string `csv:"Period"`
	Category              string `csv:"Category"`
	Cmb                   string `csv:"CMB"`
	Device                string `csv:"Device"`
	SupportedPlatforms    string `csv:"Supported Platforms"`
	ProceedsReason        string `csv:"Proceeds Reason"`
	PreservedPricing      string `csv:"Preserved Pricing"`
	Client                string `csv:"Client"`
	OrderType             string `csv:"Order Type"`
}
type AppStoreConnectAPIAccessInfo struct {
	BaseUrl     string
	AuthKeyFile string
	IssuerId    string
	KeyID       string
}
