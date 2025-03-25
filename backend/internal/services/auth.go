package services

import (
	"context"
	"errors"
	"go/token"
	"time"

	"github.com/AlexanderWangY/groupgo-backend/internal/db/sqlc"
	"github.com/AlexanderWangY/groupgo-backend/internal/generator"
	"github.com/AlexanderWangY/groupgo-backend/internal/password"
	"github.com/AlexanderWangY/groupgo-backend/internal/repository"
	"github.com/AlexanderWangY/groupgo-backend/internal/transaction"
	"github.com/bsm/ginkgo/v2/ginkgo/generators"
	"github.com/jackc/pgx/v5"
)

var (
	ErrUserExists          = errors.New("user with that email already exists")
	ErrInvalidRefreshToken = errors.New("invalid refresh token")
	ErrUserCreationFailed = errors.New("create user failed")
	ErrSessionCreationFailed = errors.New("create session failed")
	ErrRefreshTokenCreationFailed = errors.New("create refresh token failed")
)

type AuthService struct {
	userRepo        *repository.UserRepository
	sessionRepo     *repository.SessionRepository
	txManager       *transaction.TransactionManager
	sessionDuration time.Duration
	refreshDuration time.Duration
}

func NewAuthService(userRepo *repository.UserRepository, sessionRepo *repository.SessionRepository, txManager *transaction.TransactionManager) *AuthService {
	return &AuthService{
		userRepo:        userRepo,
		sessionRepo:     sessionRepo,
		txManager:       txManager,
		sessionDuration: 24 * time.Hour,      // 24 hour session length
		refreshDuration: 30 * 24 * time.Hour, // 30 day refresh length
	}
}

func (s *AuthService) Register(ctx context.Context, email, pass string) (*sqlc.AuthSession, *sqlc.AuthRefreshToken, error) {
	var authSession *sqlc.AuthSession
	var refreshToken *sqlc.AuthRefreshToken

	err := s.txManager.WithTx(ctx, func(q *sqlc.Queries, tx pgx.Tx) error {
		user, err := s.userRepo.GetByEmail(ctx, email)
		if err == nil && user != nil {
			return ErrUserExists
		}

		hashedPass, err := password.HashPassword(pass)
		if err != nil {
			return password.ErrHashFailed
		}

		newUser, err := s.userRepo.Create(ctx, sqlc.CreateUserParams{
			Email:        email,
			PasswordHash: hashedPass,
		})
		if err != nil {
			return ErrUserCreationFailed
		}

		// Generate tokens
		sessionToken, err := generator.GenerateSecureToken(32)
		if err != nil {
			return err
		}

		newRefreshToken, err := generator.GenerateSecureToken(64)
		if err != nil {
			return err
		}

		authSession, err = s.sessionRepo.CreateSession(ctx, newUser.ID, sessionToken, s.sessionDuration)
		if err != nil {
			return err
		}

		refreshToken, err = 

	})
}
