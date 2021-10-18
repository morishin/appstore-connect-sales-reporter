package reporter

import (
	"fmt"

	"github.com/slack-go/slack"
)

func postSlack(webhookUrl string, proceeds *Proceeds) {
	formatPrice := func(price int) string {
		if price < 10000 {
			return fmt.Sprintf("￥%d", price)
		} else {
			return fmt.Sprintf("￥%.2f万", float32(price)/10000)
		}
	}

	fields := []slack.AttachmentField{
		{
			Title: "2日前",
			Value: formatPrice(proceeds.DayBeforeYesterday),
			Short: true,
		}, {
			Title: "3日前",
			Value: formatPrice(proceeds.ThreeDaysAgo),
			Short: true,
		}, {
			Title: "今月",
			Value: formatPrice(proceeds.ThisMonth),
			Short: true,
		}, {
			Title: "先月",
			Value: formatPrice(proceeds.LastMonth),
			Short: true,
		},
	}
	attachments := []slack.Attachment{{
		Color:      "#0071e3",
		AuthorName: "AppStore アプリ内課金収益",
		Fields:     fields,
	}}
	message := slack.WebhookMessage{
		Username:    "appstore-connect-sales-reporter",
		Attachments: attachments,
	}
	slack.PostWebhook(webhookUrl, &message)
}
