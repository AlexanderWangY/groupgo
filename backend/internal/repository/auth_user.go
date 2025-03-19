package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/AlexanderWangY/swoppr-backend/internal/db/sqlc"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

var (
	ErrNoUserFound   = errors.New("no user found")
	ErrDuplicateUser = errors.New("user already exists")
	ErrUnknown       = errors.New("unknown error occurred")
)

type UserRepository struct {
	db *sqlc.Queries
}

func NewUserRepository(db *sqlc.Queries) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) Create(ctx context.Context, params sqlc.CreateUserParams, txs ...pgx.Tx) (*sqlc.AuthUser, error) {
	qtx := withOptionalTx(r.db, txs...)

	user, err := qtx.CreateUser(ctx, params)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *UserRepository) GetByID(ctx context.Context, id uuid.UUID) (*sqlc.AuthUser, error) {
	user, err := r.db.GetUserByID(ctx, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrNoUserFound
		}

		return nil, ErrUnknown
	}

	return &user, nil
}

func (r *UserRepository) GetByEmail(ctx context.Context, email string) (*sqlc.AuthUser, error) {
	user, err := r.db.GetUserByEmail(ctx, email)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrNoUserFound
		}

		return nil, ErrUnknown
	}

	return &user, nil
}
