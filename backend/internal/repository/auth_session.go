package repository

import (
	"context"

	"github.com/AlexanderWangY/swoppr-backend/internal/db/sqlc"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
)

type SessionRepository struct {
	db *sqlc.Queries
}

func NewSessionRepository(db *sqlc.Queries) *SessionRepository {
	return &SessionRepository{
		db: db,
	}
}

func (r *SessionRepository) CreateSessionWithTx(ctx context.Context, tx pgx.Tx, userId pgtype.UUID) (*sqlc.AuthSession, error) {
	qtx := r.db.WithTx(tx)

	session, err := qtx.CreateSession(ctx, userId)
	if err != nil {
		return nil, err
	}

	return &session, nil
}
