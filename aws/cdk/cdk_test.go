package main

import (
	"encoding/json"
	"testing"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/stretchr/testify/assert"
	"github.com/tidwall/gjson"
)

func TestCdkStack(t *testing.T) {
	// GIVEN
	app := awscdk.NewApp(nil)
	setUpEnv(t)

	// WHEN
	stack := NewCdkStack(app, "MyStack", nil)

	// THEN
	bytes, err := json.Marshal(app.Synth(nil).GetStackArtifact(stack.ArtifactId()).Template())
	if err != nil {
		t.Error(err)
	}

	template := gjson.ParseBytes(bytes)
	functionName := template.Get("Resources.appstoreconnectsalesreporterfunction55CD227D.Properties.FunctionName").String()
	cronExpression := template.Get("Resources.appstoreconnectsalesreporterrule9B4308FB.Properties.ScheduleExpression").String()
	env := template.Get("Resources.appstoreconnectsalesreporterfunction55CD227D.Properties.Environment.Variables").Map()
	assert.Equal(t, "https://api.appstoreconnect.apple.com", env["APP_STORE_CONNECT_API_BASE_URL"].String())
	assert.Equal(t, "appstore-connect-sales-reporter-function", functionName)
	assert.Equal(t, "cron(0 0 * * ? *)", cronExpression)
}

func setUpEnv(t *testing.T) {
	t.Setenv("APP_STORE_CONNECT_KEY_ID", "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx")
	t.Setenv("APP_STORE_CONNECT_ISSUER_ID", "XXXXXXXXXX")
	t.Setenv("APP_STORE_CONNECT_API_BASE_URL", "https://api.appstoreconnect.apple.com")
	t.Setenv("SLACK_WEBHOOK_URL", "https://hooks.slack.com/services/T00000000/B00000000/XXXXXXXXXXXXXXXXXXXXXXXX")
	t.Setenv("CURRENCY", "JPY")
	t.Setenv("CRON", "cron(0 0 * * ? *)")
	t.Setenv("TZ", "UTC")
}
