package env

import (
	"errors"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func Load() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("No .env file found, please create and populate a .env in root")
	}
}

func Get(key string) (*string, error) {
	value := os.Getenv(key)
	if value == "" {
		return nil, errors.New("Key not found in .env")
	}

	return &value, nil
}
