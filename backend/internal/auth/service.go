package auth

import (
	"context"
	"time"

	"github.com/AlexanderWangY/swoppr-backend/db"
	"github.com/AlexanderWangY/swoppr-backend/db/sqlc"
	"github.com/AlexanderWangY/swoppr-backend/internal/utils"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

type UserService struct {
	db *db.Database
}

func NewUserService(database *db.Database) *UserService {
	return &UserService{
		db: database,
	}
}

func (s *UserService) CreateUser(user sqlc.CreateUserParams) error {
	_, err := s.db.Query.CreateUser(context.Background(), user)
	return err
}

func (s *UserService) GetUser(userId uuid.UUID) (sqlc.AuthUser, error) {
	return s.db.Query.GetUserByID(context.Background(), userId)
}

func (s *UserService) RegisterUserAndJWT(user sqlc.CreateUserParams) (*AuthResponse, error) {
	tx, err := s.db.Pool.Begin(context.Background())
	if err != nil {
		return nil, err
	}
	defer tx.Rollback(context.Background())

	qtx := s.db.Query.WithTx(tx)

	userId, err := qtx.CreateUser(context.Background(), user)
	if err != nil {
		return nil, err
	}

	session, err := qtx.CreateSession(context.Background(), pgtype.UUID{
		Bytes: userId,
		Valid: true,
	})
	if err != nil {
		return nil, err
	}

	accessToken, err := s.GenerateAndStoreAccessToken(qtx, userId, session.ID)
	if err != nil {
		return nil, err
	}
	refreshToken, err := s.GenerateAndStoreRefreshToken(qtx, userId, session.ID)
	if err != nil {
		return nil, err
	}

	err = tx.Commit(context.Background())
	if err != nil {
		return nil, err
	}

	res := AuthResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	return &res, nil

}

func (s *UserService) GenerateAndStoreAccessToken(qtx *sqlc.Queries, userId, sessionId uuid.UUID) (string, error) {
	accessTokenExpTime := time.Now().Add(24 * time.Hour)
	accessTokenId := uuid.New()

	accessToken, err := utils.GenerateJWT(userId.String(), accessTokenId.String(), accessTokenExpTime)
	if err != nil {
		return "", err
	}

	_, err = qtx.CreateAccessToken(context.Background(), sqlc.CreateAccessTokenParams{
		ID:        accessTokenId,
		Token:     accessToken,
		SessionID: pgtype.UUID{Bytes: sessionId, Valid: true},
		ExpiresAt: pgtype.Timestamptz{
			Time:  accessTokenExpTime,
			Valid: true,
		},
	})
	if err != nil {
		return "", err
	}

	return accessToken, nil
}

func (s *UserService) GenerateAndStoreRefreshToken(qtx *sqlc.Queries, userId, sessionId uuid.UUID) (string, error) {
	refreshTokenExpTime := time.Now().AddDate(0, 1, 0)

	refreshToken, err := utils.GenerateRefreshToken()
	if err != nil {
		return "", err
	}

	_, err = qtx.CreateRefreshToken(context.Background(), sqlc.CreateRefreshTokenParams{
		Token:     refreshToken,
		SessionID: pgtype.UUID{Bytes: sessionId, Valid: true},
		ExpiresAt: pgtype.Timestamptz{
			Time:  refreshTokenExpTime,
			Valid: true,
		},
	})
	if err != nil {
		return "", err
	}

	return refreshToken, nil
}
