package seeder

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/jgndev/horsefacts-api/pkg/constants"
	"github.com/jgndev/horsefacts-api/seed/reader"
	"github.com/spf13/viper"
	"log"
	"os"
	"path/filepath"
)

const (
	FactsFile  = "json/facts.json"
	BreedsFile = "json/breeds.json"
)

func SeedHorseFacts(db *dynamodb.DynamoDB) {
	factsPath := filepath.Join(filepath.Dir(os.Args[0]), FactsFile)
	facts, err := reader.ReadFactsFromJSON(factsPath)
	if err != nil {
		log.Fatalf("Failed to read HorseFacts from JSON: %s", err.Error())
	}

	for _, fact := range facts {
		av, err := dynamodbattribute.MarshalMap(fact)
		if err != nil {
			log.Fatalf("Caught error marshalling new fact item: %s", err.Error())
		}

		input := &dynamodb.PutItemInput{
			Item:      av,
			TableName: aws.String(viper.GetString(constants.FactsTable)),
		}

		_, err = db.PutItem(input)
		if err != nil {
			log.Fatalf("Caught error calling PutItem for fact: %s", err.Error())
		}
	}

	log.Println("Completed seeded HorseFacts table")
}

func SeedHorseBreeds(db *dynamodb.DynamoDB) {
	breedsPath := filepath.Join(filepath.Dir(os.Args[0]), BreedsFile)
	breeds, err := reader.ReadBreedsFromJSON(breedsPath)
	if err != nil {
		log.Fatalf("Failed to read HorseBreeds from JSON: %s", err.Error())
	}

	for _, breed := range breeds {
		av, err := dynamodbattribute.MarshalMap(breed)
		if err != nil {
			log.Fatalf("Caught error marshalling new bree item: %s", err.Error())
		}

		input := &dynamodb.PutItemInput{
			Item:      av,
			TableName: aws.String(viper.GetString(constants.BreedsTable)),
		}

		_, err = db.PutItem(input)
		if err != nil {
			log.Fatalf("Caught error calling PutItem for breed: %s", err.Error())
		}
	}

	log.Println("Completed seeding HorseBreeds table")
}
