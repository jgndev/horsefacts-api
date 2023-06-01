package handler

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/gofiber/fiber/v2"
	"github.com/jgndev/horsefacts-api/pkg/config"
	"github.com/jgndev/horsefacts-api/pkg/database"
	"github.com/jgndev/horsefacts-api/pkg/types"
	"log"
)

func GetHealthHandler(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "ok",
	})
}

func GetFactHandler(c *fiber.Ctx) error {

	cfg, err := config.GetConfig()
	if err != nil {
		log.Fatalf("Failed to read required configuration: %v", err.Error())
	}

	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(cfg.AWSRegion),
		Credentials: credentials.NewStaticCredentials(
			cfg.AWSCredentials.ClientID,
			cfg.AWSCredentials.ClientSecret,
			""),
	})

	if err != nil {
		log.Printf("Failed to create AWS session: %v", err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create AWS session",
		})
	}

	db := dynamodb.New(sess)

	fact, err := database.GetRandomFact(db)
	if err != nil {
		log.Printf("Failed to get random fact: %v", err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to get random fact",
		})
	}

	return c.JSON(types.FactResponse{Fact: fact.Fact})
}

func GetBreedHandler(c *fiber.Ctx) error {
	cfg, err := config.GetConfig()
	if err != nil {
		log.Fatalf("Failed to read required configuration: %v", err.Error())
	}

	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(cfg.AWSRegion),
		Credentials: credentials.NewStaticCredentials(
			cfg.AWSCredentials.ClientID,
			cfg.AWSCredentials.ClientSecret,
			""),
	})

	if err != nil {
		log.Printf("Failed to create AWS session: %v", err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create AWS session",
		})
	}

	db := dynamodb.New(sess)

	breed, err := database.GetRandomBreed(db)
	if err != nil {
		log.Printf("Failed to get random breed: %v", err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to get random breed",
		})
	}

	breedResponse := types.BreedResponse{
		Name:        breed.Name,
		Country:     breed.Country,
		Colors:      breed.Colors,
		Established: breed.Established,
	}

	return c.JSON(breedResponse)
}

func GetBreedByIdHandler(c *fiber.Ctx) error {
	cfg, err := config.GetConfig()
	if err != nil {
		log.Fatalf("Failed to read required configuration: %v", err.Error())
	}

	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(cfg.AWSRegion),
		Credentials: credentials.NewStaticCredentials(
			cfg.AWSCredentials.ClientID,
			cfg.AWSCredentials.ClientSecret,
			""),
	})

	if err != nil {
		log.Printf("Failed to create AWS session: %v", err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create AWS session",
		})
	}

	db := dynamodb.New(sess)

	breedId := c.Params("id")

	breed, err := database.GetBreedById(db, breedId)
	if err != nil {
		log.Printf("Failed to get breed with ID: %v", breedId)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Failed to get breed with specified ID, bad request",
		})
	}

	breedResponse := types.BreedResponse{
		Name:        breed.Name,
		Country:     breed.Country,
		Colors:      breed.Colors,
		Established: breed.Established,
	}

	return c.JSON(breedResponse)
}
