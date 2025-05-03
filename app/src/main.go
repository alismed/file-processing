package main

import (
	"context"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type Order struct {
	id string
	name float64
	region string
	description string
}

var (
	s3Client *s3.Client
)

func handler(_ context.Context, event map[string]interface{}) (string, error) {
	return "Hi lambda", nil
}

func main() {
	lambda.Start(handler)
}