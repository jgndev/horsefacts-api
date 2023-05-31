package config

import "os"

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
