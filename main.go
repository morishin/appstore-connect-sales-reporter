package main

import (
	"os"

	"github.com/joho/godotenv"
)

func main() {
	accessInfo, slackWebhookUrl, currency := loadEnv()
	salesReports := getSalesReports(accessInfo)
	proceeds := salesReportsToProceeds(&salesReports, currency)
	postSlack(slackWebhookUrl, &proceeds)
}

func loadEnv() (AppStoreConnectAPIAccessInfo, string, string) {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	accessInfo := AppStoreConnectAPIAccessInfo{
		BaseUrl:     os.Getenv("APP_STORE_CONNECT_API_BASE_URL"),
		AuthKeyFile: os.Getenv("APP_STORE_CONNECT_AUTH_KEY_P8_FILE_PATH"),
		IssuerId:    os.Getenv("APP_STORE_CONNECT_ISSUER_ID"),
		KeyID:       os.Getenv("APP_STORE_CONNECT_KEY_ID"),
	}
	slackWebhookUrl := os.Getenv("SLACK_WEBHOOK_URL")
	currency := os.Getenv("CURRENCY")
	return accessInfo, slackWebhookUrl, currency
}
