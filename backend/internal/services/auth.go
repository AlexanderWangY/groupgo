package services

import (
	"context"
	"errors"
	"time"

	"github.com/AlexanderWangY/swoppr-backend/internal/db/sqlc"
	"github.com/AlexanderWangY/swoppr-backend/internal/repository"
)

var (
	ErrUserExists          = errors.New("user with that email already exists")
	ErrInvalidRefreshToken = errors.New("invalid refresh token")
)

type AuthService struct {
	userRepo        *repository.UserRepository
	sessionRepo     *repository.SessionRepository
	sessionDuration time.Duration
	refreshDuration time.Duration
}

func NewAuthService(userRepo *repository.UserRepository) *AuthService {
	return &AuthService{
		userRepo:        userRepo,
		sessionDuration: 24 * time.Hour,      // 24 hour session length
		refreshDuration: 30 * 24 * time.Hour, // 30 day refresh length
	}
}

func (s *AuthService) Register(ctx context.Context, email, password string) (sqlc.AuthSession, sqlc.AuthRefreshToken, error) {

}
