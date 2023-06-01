package database

import (
	"errors"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/jgndev/horsefacts-api/pkg/config"
	"github.com/jgndev/horsefacts-api/pkg/interfaces"
	"github.com/jgndev/horsefacts-api/pkg/types"
	"log"
	"math/rand"
	"strconv"
)

func GetRandomFact(db interfaces.DynamoDbInterface) (*types.Fact, error) {
	cfg, err := config.GetConfig()
	if err != nil {
		log.Fatalf("Failed to read required configuration: %v", err.Error())
	}

	input := &dynamodb.ScanInput{
		TableName: aws.String(cfg.FactsTable),
	}

	result, err := db.Scan(input)
	if err != nil {
		log.Printf("Error scanning %v: %v", cfg.FactsTable, err.Error())
		return nil, err
	}

	if len(result.Items) == 0 {
		return nil, errors.New("no facts found in the table")
	}

	randomIndex := rand.Intn(len(result.Items))

	factItem := result.Items[randomIndex]

	var fact types.Fact
	if err = dynamodbattribute.UnmarshalMap(factItem, &fact); err != nil {
		log.Printf("Error unmarshalling fact item: %v", err.Error())
		return nil, err
	}

	return &fact, nil
}

func GetRandomBreed(db interfaces.DynamoDbInterface) (*types.Breed, error) {
	cfg, err := config.GetConfig()
	if err != nil {
		log.Fatalf("Failed to read required configuration: %v", err.Error())
	}

	input := &dynamodb.ScanInput{
		TableName: aws.String(cfg.BreedsTable),
	}

	result, err := db.Scan(input)
	if err != nil {
		log.Printf("Error scanning %v: %v", cfg.BreedsTable, err.Error())
		return nil, err
	}

	if len(result.Items) == 0 {
		return nil, errors.New("no breeds found in the table")
	}

	randomIndex := rand.Intn(len(result.Items))

	breedItem := result.Items[randomIndex]

	var breed types.Breed
	if err = dynamodbattribute.UnmarshalMap(breedItem, &breed); err != nil {
		log.Printf("Error unmarshalling breed item: %v", err.Error())
		return nil, err
	}

	return &breed, nil
}

func GetBreedById(db interfaces.DynamoDbInterface, id string) (*types.Breed, error) {
	cfg, err := config.GetConfig()
	if err != nil {
		log.Fatalf("Failed to read required configuration: %v", err.Error())
	}

	idInt, err := strconv.Atoi(id)
	if err != nil {
		log.Printf("Error converting id to int: %v", err.Error())
		return nil, err
	}

	input := &dynamodb.GetItemInput{
		TableName: aws.String(cfg.BreedsTable),
		Key: map[string]*dynamodb.AttributeValue{
			"ID": {
				N: aws.String(strconv.Itoa(idInt)),
			},
		},
	}

	result, err := db.GetItem(input)
	if err != nil {
		log.Printf("Error getting breed with ID %v: %v", id, err.Error())
		return nil, err
	}

	if result.Item == nil {
		return nil, errors.New("no breed found with the given id")
	}

	var breed types.Breed
	if err = dynamodbattribute.UnmarshalMap(result.Item, &breed); err != nil {
		log.Printf("Error unmarshalling breed item: %v", err.Error())
		return nil, err
	}

	return &breed, nil
}
