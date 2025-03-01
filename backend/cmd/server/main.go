package main

import (
	"log"
	"os"

	"github.com/AlexanderWangY/swoppr-backend/api"
	"github.com/AlexanderWangY/swoppr-backend/db"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Something went wrong loading the .env")
	}

	connection_url := os.Getenv("DATABASE_URL")
	if connection_url == "" {
		log.Fatal("Could not find DATABASE_URL in .env")
	}

	new_db := db.NewDatabase(connection_url)
	defer new_db.Pool.Close()

	api.StartServer(new_db)
}
