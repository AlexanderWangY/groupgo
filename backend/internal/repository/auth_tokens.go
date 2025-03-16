package repository

import "github.com/AlexanderWangY/swoppr-backend/internal/db/sqlc"

type TokenRepository struct {
	db *sqlc.Queries
}

func NewTokenRepository(db *sqlc.Queries) *TokenRepository {
	return &TokenRepository{
		db: db,
	}
}
