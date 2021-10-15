package reporter

import (
	"os"

	"github.com/joho/godotenv"
)

func Run() {
	accessInfo, slackWebhookUrl, currency := loadEnv()
	salesReports := getSalesReports(accessInfo)
	proceeds := salesReportsToProceeds(&salesReports, currency)
	postSlack(slackWebhookUrl, &proceeds)
}

func loadEnv() (*AppStoreConnectAPIAccessInfo, string, string) {
	godotenv.Load()
	accessInfo := AppStoreConnectAPIAccessInfo{
		BaseUrl:  os.Getenv("APP_STORE_CONNECT_API_BASE_URL"),
		IssuerId: os.Getenv("APP_STORE_CONNECT_ISSUER_ID"),
		KeyID:    os.Getenv("APP_STORE_CONNECT_KEY_ID"),
	}
	slackWebhookUrl := os.Getenv("SLACK_WEBHOOK_URL")
	currency := os.Getenv("CURRENCY")
	return &accessInfo, slackWebhookUrl, currency
}
