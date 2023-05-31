package interfaces

import "github.com/aws/aws-sdk-go/service/dynamodb"

type DynamoDbApi interface {
	Scan(input *dynamodb.ScanInput) (*dynamodb.ScanOutput, error)
	GetItem(input *dynamodb.GetItemInput) (*dynamodb.GetItemOutput, error)
}
