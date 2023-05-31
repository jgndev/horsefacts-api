package config

import (
	"log"
	"os"
)

type Config struct {
	Dev            bool
	FactsTable     string
	BreedsTable    string
	AWSRegion      string
	AWSCredentials struct {
		ClientID     string
		ClientSecret string
	}
}

func GetConfig() (*Config, error) {
	var cfg Config

	cfg.Dev = true
	cfg.FactsTable = os.Getenv("FACTS_TABLE")
	cfg.BreedsTable = os.Getenv("BREEDS_TABLE")
	cfg.AWSRegion = os.Getenv("AWS_REGION")
	cfg.AWSCredentials.ClientID = os.Getenv("AWS_CLIENT_ID")
	cfg.AWSCredentials.ClientSecret = os.Getenv("AWS_CLIENT_SECRET")

	return &cfg, nil
}

func PrintConfigStatus() {
	log.Println("Checking state of required environment variables")

	cfg, err := GetConfig()
	if err != nil {
		log.Printf("Error loading configuration: %v", err.Error())
	}

	if cfg.FactsTable != "" {
		log.Printf("%-20s: %t\n", "Facts Table set", true)
	} else {
		log.Printf("%-20s: %t\n", "Facts Table set", false)
	}

	if cfg.BreedsTable != "" {
		log.Printf("%-20s: %t\n", "Breeds Table set", true)
	} else {
		log.Printf("%-20s: %t\n", "Breeds Table set", false)
	}

	if cfg.AWSRegion != "" {
		log.Printf("%-20s: %t\n", "AWS Region set", true)
	} else {
		log.Printf("%-20s: %t\n", "AWS Region set", false)
	}

	if cfg.AWSCredentials.ClientID != "" {
		log.Printf("%-20s: %t\n", "Client ID set", true)
	} else {
		log.Printf("%-20s: %t\n", "Client ID set", false)
	}

	if cfg.AWSCredentials.ClientSecret != "" {
		log.Printf("%-20s: %t\n", "Client Secret set", true)
	} else {
		log.Printf("%-20s: %t\n", "Client Secret set", false)
	}
}
