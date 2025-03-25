package generator

import (
	"crypto/rand"
	"encoding/base64"
	"errors"
)

var (
	ErrFailedToGenerateToken = errors.New("failed to generate token")
)

func GenerateSecureToken(length int) (string, error) {
	bytes := make([]byte, length)
	_, err := rand.Read(bytes)
	if err != nil {
		return "", ErrFailedToGenerateToken
	}
	return base64.URLEncoding.EncodeToString(bytes), nil
}
