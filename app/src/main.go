package main

import (
	"bufio"
	"context"
	"encoding/csv"
	"encoding/json"
	"log"
	"os"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	ddbTypes "github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type Order struct {
	Id          string
	Region      string
	Name        string
	Description string
}

type S3Event struct {
	Records []struct {
		S3 struct {
			Bucket struct {
				Name string `json:"name"`
			} `json:"bucket"`
			Object struct {
				Key string `json:"key"`
			} `json:"object"`
		} `json:"s3"`
	} `json:"Records"`
}

func handler(ctx context.Context, event json.RawMessage) (string, error) {
	log.Printf("Event received: %s", string(event))

	var s3Event S3Event
	if err := json.Unmarshal(event, &s3Event); err != nil {
		log.Printf("Failed to unmarshal event: %v", err)
		return "", err
	}
	if len(s3Event.Records) == 0 {
		log.Printf("No records in event")
		return "", nil
	}
	bucket := s3Event.Records[0].S3.Bucket.Name
	key := s3Event.Records[0].S3.Object.Key
	log.Printf("Bucket: %s, Key: %s", bucket, key)

	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		log.Printf("Failed to load AWS config: %v", err)
		return "", err
	}
	s3Client := s3.NewFromConfig(cfg)
	ddbClient := dynamodb.NewFromConfig(cfg)

	// Baixar o arquivo do S3
	getObj, err := s3Client.GetObject(ctx, &s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	})
	if err != nil {
		log.Printf("Failed to get object from S3: %v", err)
		return "", err
	}
	defer getObj.Body.Close()

	r := csv.NewReader(bufio.NewReader(getObj.Body))
	r.Comma = ';'
	r.FieldsPerRecord = -1

	_, err = r.Read()
	if err != nil {
		log.Printf("Failed to read CSV header: %v", err)
		return "", err
	}

	for {
		record, err := r.Read()
		if err != nil {
			break // EOF
		}
		order := Order{
			Id:          record[0],
			Region:      record[1],
			Name:        record[2],
			Description: record[3],
		}
		item := map[string]ddbTypes.AttributeValue{
			"Id":          &ddbTypes.AttributeValueMemberS{Value: order.Id},
			"Region":      &ddbTypes.AttributeValueMemberS{Value: order.Region},
			"Name":        &ddbTypes.AttributeValueMemberS{Value: order.Name},
			"Description": &ddbTypes.AttributeValueMemberS{Value: order.Description},
		}
		log.Printf("Item to insert: %+v", item)
		log.Printf("Item to insert: %s, %s, %s, %s", order.Id, order.Region, order.Name, order.Description)
		log.Printf("Table name: %s", os.Getenv("DYNAMODB_TABLE"))
		_, err = ddbClient.PutItem(ctx, &dynamodb.PutItemInput{
			TableName: aws.String(os.Getenv("DYNAMODB_TABLE")),
			Item:      item,
		})
		
		if err != nil {
			log.Printf("Failed to put item in DynamoDB: %v", err)
		}
	}
	return "Processamento concluído", nil
}

func main() {
	lambda.Start(handler)
}