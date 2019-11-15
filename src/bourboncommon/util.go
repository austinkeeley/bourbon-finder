package bourboncommon

import (
	"encoding/json"
	"log"
	"os"
)

type Store struct {
	Name    string `json:"name"`
	StoreID int    `json:"storeID"`
}

type Product struct {
	Name        string `json:"name"`
	ProductCode string `json:"productCode"`
}

type Config struct {
	Stores   []Store   `json:"stores"`
	Wishlist []Product `json:"wishlist"`
}

func OpenConfig(fileName string) (*Config, error) {
	log.Println("Opening file " + fileName)
	file, err := os.Open(fileName)
	if err != nil {
		log.Println("Error opening file")
		return nil, err
	}
	defer file.Close()

	var config *Config
	err = json.NewDecoder(file).Decode(&config)
	return config, err
}
