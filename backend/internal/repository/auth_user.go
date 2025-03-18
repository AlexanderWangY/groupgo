package repository

import (
	"github.com/AlexanderWangY/swoppr-backend/internal/db/sqlc"
)

type UserRepository struct {
	db *sqlc.Queries
}

func NewUserRepository(db *sqlc.Queries) *UserRepository {
	return &UserRepository{
		db: db,
	}
}
