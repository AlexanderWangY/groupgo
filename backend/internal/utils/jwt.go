package utils

import (
	"crypto/rand"
	"encoding/base64"
	"errors"
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var JWTSecret = []byte(os.Getenv("JWT_SECRET"))

func GenerateJWT(userId, tokenId string, exp_time time.Time) (string, error) {
	claims := jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(exp_time),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		NotBefore: jwt.NewNumericDate(time.Now()),
		Issuer:    "swoppr-backend",
		Subject:   userId,
		ID:        tokenId,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	if JWTSecret == nil {
		log.Println("No JWT_SECRET present, failing at GenerateJWT in internals/utils/jwt.go")
		return "", errors.New("No JWT_Secret present")
	}

	tokenString, err := token.SignedString(JWTSecret)
	if err != nil {
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

// func validToken(token string) *error {
// 	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
// 		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
// 			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
// 		}

// 		return JWTSecret, nil
// 	})

// 	if err != nil {
// 		return &err
// 	}

// 	if claims, ok := parsedToken.Claims.(jwt.MapClaims); !ok {
// 		return nil
// 	}

// }

// Helper function to generate random string for refresh token
func GenerateRandomString(length int) (string, error) {
	b := make([]byte, length)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}

	return base64.URLEncoding.EncodeToString(b), nil
}
