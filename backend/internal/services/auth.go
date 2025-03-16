package services

import (
	"context"

	"github.com/AlexanderWangY/swoppr-backend/internal/creds"
	"github.com/AlexanderWangY/swoppr-backend/internal/db"
	"github.com/AlexanderWangY/swoppr-backend/internal/db/sqlc"
	"github.com/AlexanderWangY/swoppr-backend/internal/models"
	"github.com/AlexanderWangY/swoppr-backend/internal/repository"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
)

type AuthService struct {
	userRepo    *repository.UserRepository
	tokenRepo   *repository.TokenRepository
	sessionRepo *repository.SessionRepository
	db          *db.Database
}

func NewAuthService(userRepo *repository.UserRepository, tokenRepo *repository.TokenRepository, sessionRepo *repository.SessionRepository, db *db.Database) *AuthService {
	return &AuthService{
		userRepo:    userRepo,
		tokenRepo:   tokenRepo,
		sessionRepo: sessionRepo,
		db:          db,
	}
}

func (a *AuthService) Register(ctx context.Context, email, password string) (*models.AuthResponse, error) {
	// Hash password
	hashedPassword, err := creds.HashPassword(password)
	if err != nil {
		return nil, err
	}

	// Begin transaction to register user
	// Create User, Create Session, Create Access and refresh token

	tx, err := a.db.Pool.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return nil, err
	}
	defer tx.Rollback(ctx)

	// First create a user
	user, err := a.userRepo.CreateUserWithTx(ctx, tx, sqlc.CreateUserParams{
		Email:        email,
		PasswordHash: hashedPassword,
	})
	if err != nil {
		return nil, err
	}

	session, err := a.sessionRepo.CreateSessionWithTx(ctx, tx, pgtype.UUID{Bytes: user.ID, Valid: true})
	if err != nil {
		return nil, err
	}

	// TODO: Add custom claims to a string map such as roles and generate tokens

}
