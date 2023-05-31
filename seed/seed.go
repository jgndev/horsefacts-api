package main

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/jgndev/horsefacts-api/pkg/constants"
	"github.com/jgndev/horsefacts-api/seed/seeder"
	"github.com/spf13/viper"
	"log"
)

const EnvFile = ".env"

func main() {
	if constants.DEVENV {
		viper.SetConfigFile(EnvFile)
		err := viper.ReadInConfig()
		if err != nil {
			log.Fatalf("Failed to read variables from environment file. %v", err.Error())
			return
		}

		sess, err := session.NewSession(&aws.Config{
			Region: aws.String(viper.GetString(constants.AwsRegion)),
			Credentials: credentials.NewStaticCredentials(
				viper.GetString(constants.AwsAccessKey),
				viper.GetString(constants.AwsAccessSecret),
				""),
		})

		if err != nil {
			log.Fatalf("Failed to connect to AWS: %s", err.Error())
		}

		db := dynamodb.New(sess)

		log.Println("Seeding facts...")
		seeder.SeedHorseFacts(db)

		log.Println("Seeding breeds...")
		seeder.SeedHorseBreeds(db)
	} else {
		log.Println("Seeding should be run in the local development environment.")
		log.Println("No seeding action has taken place.")
	}
}
