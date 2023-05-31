package reader

import (
	"encoding/json"
	types2 "github.com/jgndev/horsefacts-api/pkg/types"
	"log"
	"os"
)

func ReadFactsFromJSON(filePath string) ([]types2.Fact, error) {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("Error opening file contents for read: %v", err.Error())
		return nil, err
	}
	defer func(file *os.File) {
		err = file.Close()
		if err != nil {
			log.Printf("Caught error closing %v: %v", filePath, err.Error())
		}
	}(file)

	var facts []types2.Fact
	err = json.NewDecoder(file).Decode(&facts)
	if err != nil {
		log.Fatalf("Error decoding file contents to JSON: %v", err.Error())
		return nil, err
	}

	return facts, nil
}

func ReadBreedsFromJSON(filePath string) ([]types2.Breed, error) {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("Error opening file contents for read: %v", err.Error())
		return nil, err
	}
	defer func(file *os.File) {
		err = file.Close()
		if err != nil {
			log.Printf("Caught error closing %v: %v", filePath, err.Error())
		}
	}(file)

	var breeds []types2.Breed
	err = json.NewDecoder(file).Decode(&breeds)
	if err != nil {
		log.Fatalf("Error decoding file contents to JSON: %v", err.Error())
		return nil, err
	}

	return breeds, nil
}
