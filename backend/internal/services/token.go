package services

import (
	"time"

	"github.com/AlexanderWangY/swoppr-backend/internal/repository"
)

type TokenService struct {
	tokenRepo          *repository.TokenRepository
	secretKey          []byte
	accessTokenExpiry  time.Duration
	refreshTokenExpiry time.Duration
}

func NewTokenService(tokenRepo *repository.TokenRepository, secretKey []byte, accessTokenExpiry, refreshTokenExpiry time.Duration) *TokenService {
	return &TokenService{
		tokenRepo:          tokenRepo,
		secretKey:          secretKey,
		accessTokenExpiry:  accessTokenExpiry,
		refreshTokenExpiry: refreshTokenExpiry,
	}
}
