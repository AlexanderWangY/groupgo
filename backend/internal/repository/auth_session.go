package repository

import (
	"context"
	"errors"
	"time"

	"github.com/AlexanderWangY/groupgo-backend/internal/db/sqlc"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

var (
	ErrCreateSessionFailed = errors.New("failed to create session")
)

type SessionRepository struct {
	db *sqlc.Queries
}

func NewSessionRepository(db *sqlc.Queries) *SessionRepository {
	return &SessionRepository{
		db: db,
	}
}

func (r *SessionRepository) CreateSession(ctx context.Context, userId uuid.UUID, token string, sessionDuration time.Duration) (*sqlc.AuthSession, error) {
	session, err := r.db.CreateSession(ctx, sqlc.CreateSessionParams{
		UserID: userId,
		Token:  token,
		ExpiresAt: pgtype.Timestamptz{
			Time:  time.Now().Add(sessionDuration),
			Valid: true,
		},
	})
	if err != nil {
		return nil, ErrCreateSessionFailed
	}

	return &session, nil
}
