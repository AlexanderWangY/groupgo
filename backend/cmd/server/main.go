package main

import (
	"context"
	"log"
	"os"
	"os/signal"

	"github.com/AlexanderWangY/swoppr-backend/internal/api"
	"github.com/AlexanderWangY/swoppr-backend/internal/db"
	"github.com/joho/godotenv"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill)
	defer cancel()

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

	api.StartServer(ctx, new_db)
}
