package repository

import (
	"context"
	"errors"

	"github.com/AlexanderWangY/groupgo-backend/internal/db/sqlc"
)

var (
	ErrFailedCreateRefreshToken = errors.New("failed to create refresh token")
)

type TokenRespository struct {
	db *sqlc.Queries
}

func NewTokenRepository(db *sqlc.Queries) *TokenRespository {
	return &TokenRespository{
		db: db,
	}
}

func (r *TokenRespository) CreateRefreshToken(ctx context.Context, params sqlc.CreateRefreshTokenParams) (*sqlc.AuthRefreshToken, error) {
	refreshToken, err := r.db.CreateRefreshToken(ctx, params)
	if err != nil {
		return nil, ErrFailedCreateRefreshToken
	}
	return &refreshToken, nil
}
