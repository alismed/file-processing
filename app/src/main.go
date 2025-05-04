package main

import (
	"context"
	"encoding/json"
	"log"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type CloudTrailEvent struct {
	Detail struct {
		EventName          string `json:"eventName"`
		RequestParameters struct {
			BucketName string `json:"bucketName"`
			Key        string `json:"key"`
		} `json:"requestParameters"`
	} `json:"detail"`
}

var (
	s3Client *s3.Client
)

func handler(ctx context.Context, event map[string]interface{}) (string, error) {
	// Log the raw event for debugging
	rawEvent, _ := json.MarshalIndent(event, "", "  ")
	log.Printf("Raw event received: %s", string(rawEvent))

	// Parse the CloudTrail event
	eventBytes, err := json.Marshal(event)
	if err != nil {
		log.Printf("Error marshaling event: %v", err)
		return "", err
	}

	var cloudTrailEvent CloudTrailEvent
	if err := json.Unmarshal(eventBytes, &cloudTrailEvent); err != nil {
		log.Printf("Error unmarshaling event: %v", err)
		return "", err
	}

	// Log the processed event details
	log.Printf("Event Name: %s", cloudTrailEvent.Detail.EventName)
	log.Printf("Bucket Name: %s", cloudTrailEvent.Detail.RequestParameters.BucketName)
	log.Printf("Object Key: %s", cloudTrailEvent.Detail.RequestParameters.Key)

	return "Event processed successfully", nil
}

func main() {
	log.Printf("Lambda starting up...")
	lambda.Start(handler)
}