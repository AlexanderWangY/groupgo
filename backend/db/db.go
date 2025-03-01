package db

import (
	"context"
	"log"
	"os"

	"github.com/jackc/pgx/v5"
)

var DB *pgx.Conn

func ConnectDB() {
	db_url := os.Getenv("DATABASE_URL")
	if db_url == "" {
		log.Fatal("DATABASE_URL missing")
	}

	var err error
	DB, err = pgx.Connect(context.Background(), db_url)
	if err != nil {
		log.Fatalf("Could not connect to database at %s", db_url)
	}

	log.Print("Successfully connected to database")

}

func CloseDB() {
	if DB != nil {
		DB.Close(context.Background())
	}
	log.Print("Closed database connection")
}
