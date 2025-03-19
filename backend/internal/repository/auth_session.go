package repository

import "github.com/AlexanderWangY/swoppr-backend/internal/db/sqlc"

type SessionRepository struct {
	db *sqlc.Queries
}

func NewSessionRepository(db *sqlc.Queries) *SessionRepository {
	return &SessionRepository{
		db: db,
	}
}
