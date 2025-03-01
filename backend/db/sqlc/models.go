// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0

package sqlc

import (
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

type AuthAccessToken struct {
	ID        uuid.UUID          `json:"id"`
	SessionID pgtype.UUID        `json:"session_id"`
	Token     string             `json:"token"`
	CreatedAt pgtype.Timestamptz `json:"created_at"`
	ExpiresAt pgtype.Timestamptz `json:"expires_at"`
}

type AuthRefreshToken struct {
	ID        uuid.UUID          `json:"id"`
	SessionID pgtype.UUID        `json:"session_id"`
	Token     string             `json:"token"`
	CreatedAt pgtype.Timestamptz `json:"created_at"`
	ExpiresAt pgtype.Timestamptz `json:"expires_at"`
	IsRevoked pgtype.Bool        `json:"is_revoked"`
}

type AuthSession struct {
	ID        uuid.UUID          `json:"id"`
	UserID    pgtype.UUID        `json:"user_id"`
	CreatedAt pgtype.Timestamptz `json:"created_at"`
	UpdatedAt pgtype.Timestamptz `json:"updated_at"`
}

type AuthUser struct {
	ID              uuid.UUID          `json:"id"`
	Email           string             `json:"email"`
	PasswordHash    string             `json:"password_hash"`
	FirstName       pgtype.Text        `json:"first_name"`
	LastName        pgtype.Text        `json:"last_name"`
	IsEmailVerified bool               `json:"is_email_verified"`
	CreatedAt       pgtype.Timestamptz `json:"created_at"`
	UpdatedAt       pgtype.Timestamptz `json:"updated_at"`
}
