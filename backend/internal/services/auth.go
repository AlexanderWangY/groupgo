package services

import (
	"context"
	"errors"
	"time"

	"github.com/AlexanderWangY/groupgo-backend/internal/db/sqlc"
	"github.com/AlexanderWangY/groupgo-backend/internal/password"
	"github.com/AlexanderWangY/groupgo-backend/internal/repository"
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

func NewAuthService(userRepo *repository.UserRepository, sessionRepo *repository.SessionRepository) *AuthService {
	return &AuthService{
		userRepo: userRepo,
		sessionRepo: sessionRepo,
		sessionDuration: 24 * time.Hour,      // 24 hour session length
		refreshDuration: 30 * 24 * time.Hour, // 30 day refresh length
	}
}

func (s *AuthService) Register(ctx context.Context, email, pass string) (*sqlc.AuthSession, *sqlc.AuthRefreshToken, error) {
	// Check if user exists already
	user, err := s.userRepo.GetByEmail(ctx, email)
	if err == nil && user != nil {
		return nil, nil, ErrUserExists
	}

	// At this point, the user does not exist yet
	hashedPass, err := password.HashPassword(pass)
	if err != nil {
		return nil, nil, password.ErrHashFailed
	}

	tx, err :=
}
