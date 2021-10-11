package main

import (
	"context"

	"github.com/aws/aws-lambda-go/lambda"
	reporter "github.com/morishin/appstore-connect-sales-reporter"
)

func HandleRequest(ctx context.Context) (string, error) {
	reporter.Run()
	return "Success!", nil
}

func main() {
	lambda.Start(HandleRequest)
}
