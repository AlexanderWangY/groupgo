package db

import (
	"context"
	"log"

	"github.com/AlexanderWangY/groupgo-backend/internal/db/sqlc"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Database struct {
	Pool  *pgxpool.Pool
	Query *sqlc.Queries
}

func NewDatabase(connStr string) *Database {

	pool, err := pgxpool.New(context.Background(), connStr)
	if err != nil {
		log.Fatal(err)
	}

	query := sqlc.New(pool)

	database := Database{
		Pool:  pool,
		Query: query,
	}

	return &database
}
