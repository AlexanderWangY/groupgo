package utils

import (
	"crypto/rand"
	"encoding/base64"
	"errors"
	"log"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

var JWTSecret = []byte("1j23lkjlji1j2kkj3k2j3k3k")

func GenerateJWT(userId string, exp_time time.Time) (string, error) {
	claims := jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(exp_time),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		NotBefore: jwt.NewNumericDate(time.Now()),
		Issuer:    "swoppr-backend",
		Subject:   userId,
		ID:        uuid.NewString(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	if JWTSecret == nil {
		log.Println("No JWT_SECRET present, failing at GenerateJWT in internals/utils/jwt.go")
		return "", errors.New("No JWT_Secret present")
	}

	tokenString, err := token.SignedString(JWTSecret)
	if err != nil {
		log.Println("here 1")
		return "", err
	}

	return tokenString, nil
}

func GenerateRefreshToken() (string, error) {
	token, err := GenerateRandomString(64)
	if err != nil {
		return "", nil
	}

	return token, nil
}

// Helper function to generate random string for refresh token
func GenerateRandomString(length int) (string, error) {
	b := make([]byte, length)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}

	return base64.URLEncoding.EncodeToString(b), nil
}
