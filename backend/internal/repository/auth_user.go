package repository

import (
	"context"

	"github.com/AlexanderWangY/swoppr-backend/internal/db/sqlc"
	"github.com/google/uuid"
)

type UserRepository struct {
	db *sqlc.Queries
}

func NewUserRepository(db *sqlc.Queries) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) CreateUser(ctx context.Context, params sqlc.CreateUserParams) (*sqlc.AuthUser, error) {
	user, err := r.db.CreateUser(ctx, params)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *UserRepository) GetUserByID(ctx context.Context, userId uuid.UUID) (*sqlc.AuthUser, error) {
	user, err := r.db.GetUserByID(ctx, userId)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
