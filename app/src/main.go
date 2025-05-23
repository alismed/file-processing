package main

import (
	"context"
	"encoding/csv"
	"encoding/json"
	"io"
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

func ProcessCSV(r io.Reader, putItemFunc func(Order) error) error {
    csvReader := csv.NewReader(r)
    csvReader.Comma = ';'
    csvReader.FieldsPerRecord = -1

    _, err := csvReader.Read() // descarta o cabeçalho
    if err != nil {
        return err
    }

    for {
        record, err := csvReader.Read()
        if err != nil {
            break // EOF
        }
		if len(record) < 4 {
			log.Printf("Row ignored because it has less than 4 columns: %+v", record)
			continue
		}
		order := Order{
            Id:          record[0],
            Region:      record[1],
            Name:        record[2],
            Description: record[3],
        }
        if err := putItemFunc(order); err != nil {
            return err
        }
    }
    return nil
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

	putItemFunc := func(order Order) error {
        item := map[string]ddbTypes.AttributeValue{
            "Id":          &ddbTypes.AttributeValueMemberS{Value: order.Id},
            "Region":      &ddbTypes.AttributeValueMemberS{Value: order.Region},
            "Name":        &ddbTypes.AttributeValueMemberS{Value: order.Name},
            "Description": &ddbTypes.AttributeValueMemberS{Value: order.Description},
        }
		log.Printf("Table name: %s", os.Getenv("DYNAMODB_TABLE"))
        log.Printf("Item to insert: %s, %s, %s, %s", order.Id, order.Region, order.Name, order.Description)
        _, err := ddbClient.PutItem(ctx, &dynamodb.PutItemInput{
            TableName: aws.String(os.Getenv("DYNAMODB_TABLE")),
            Item:      item,
        })
        if err != nil {
            log.Printf("Failed to put item in DynamoDB: %v", err)
        }
        return err
    }

	err = ProcessCSV(getObj.Body, putItemFunc)
    if err != nil {
        log.Printf("Error processing CSV: %v", err)
        return "", err
    }
	return "Processing completed", nil
}

func main() {
	lambda.Start(handler)
}