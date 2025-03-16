package repository

import (
	"context"

	"github.com/AlexanderWangY/swoppr-backend/internal/db/sqlc"
	"github.com/jackc/pgx/v5"
)

type TokenRepository struct {
	db *sqlc.Queries
}

func NewTokenRepository(db *sqlc.Queries) *TokenRepository {
	return &TokenRepository{
		db: db,
	}
}

func (r *TokenRepository) CreateAccessTokenWithTx(ctx context.Context, tx pgx.Tx, params sqlc.CreateAccessTokenParams) (*sqlc.AuthAccessToken, error) {
	qtx := r.db.WithTx(tx)

	accessToken, err := qtx.CreateAccessToken(ctx, params)
	if err != nil {
		return nil, err
	}

	return &accessToken, nil
}

func (r *TokenRepository) CreateRefreshToken(ctx context.Context, tx pgx.Tx, params sqlc.CreateRefreshTokenParams) (*sqlc.AuthRefreshToken, error) {
	qtx := r.db.WithTx(tx)

	refreshToken, err := qtx.CreateRefreshToken(ctx, params)
	if err != nil {
		return nil, err
	}

	return &refreshToken, nil
}
