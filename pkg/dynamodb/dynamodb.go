package dynamodb

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/jgndev/horsefacts-api/pkg/constants"
	"github.com/jgndev/horsefacts-api/pkg/types"
	"log"
	"math/rand"
	"time"
)

const FactsTable = "HorseFacts"

func GetRandomFact(db *dynamodb.DynamoDB) (*types.Fact, error) {
	input := &dynamodb.ScanInput{
		TableName: aws.String(FactsTable),
	}

	result, err := db.Scan(input)
	if err != nil {
		log.Printf("Error scanning %v: %v", constants.FactsTable, err.Error())
		return nil, err
	}

	// Generate a random index within the range of available facts
	rand.Seed(time.Now().Unix())
	randomIndex := rand.Intn(len(result.Items))

	factItem := result.Items[randomIndex]

	var fact types.Fact
	err = dynamodbattribute.UnmarshalMap(factItem, &fact)
	if err != nil {
		log.Printf("Error unmarshalling fact item: %v", err.Error())
		return nil, err
	}

	return &fact, nil
}
