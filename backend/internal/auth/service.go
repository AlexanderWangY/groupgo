package auth

import (
	"context"
	"log"
	"time"

	"github.com/AlexanderWangY/swoppr-backend/db"
	"github.com/AlexanderWangY/swoppr-backend/db/sqlc"
	"github.com/AlexanderWangY/swoppr-backend/internal/utils"
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

func (s *UserService) RegisterUserAndJWT(user sqlc.CreateUserParams) (*AuthResponse, error) {
	tx, err := s.db.Pool.Begin(context.Background())
	if err != nil {
		return nil, err
	}
	defer tx.Rollback(context.Background())

	qtx := s.db.Query.WithTx(tx)

	user_id, err := qtx.CreateUser(context.Background(), user)
	if err != nil {
		log.Println("here 2")
		return nil, err
	}

	access_token_exp_time := time.Now().Add(24 * time.Hour)
	refresh_token_exp_time := time.Now().AddDate(0, 1, 0)

	access_token, err := utils.GenerateJWT(user_id.String(), access_token_exp_time)
	if err != nil {
		log.Println("here 3")
		return nil, err
	}

	refresh_token, err := utils.GenerateRefreshToken()
	if err != nil {
		log.Println("here 4")
		return nil, err
	}

	session, err := qtx.CreateSession(context.Background(), pgtype.UUID{
		Bytes: user_id,
		Valid: true,
	})
	if err != nil {
		log.Println("here 5")
		return nil, err
	}

	_, err = qtx.CreateAccessToken(context.Background(), sqlc.CreateAccessTokenParams{
		Token:     access_token,
		SessionID: pgtype.UUID{Bytes: session.ID, Valid: true},
		ExpiresAt: pgtype.Timestamptz{
			Time:  access_token_exp_time,
			Valid: true,
		},
	})
	if err != nil {
		log.Println("here 6")
		return nil, err
	}

	_, err = qtx.CreateRefreshToken(context.Background(), sqlc.CreateRefreshTokenParams{
		Token:     refresh_token,
		SessionID: pgtype.UUID{Bytes: session.ID, Valid: true},
		ExpiresAt: pgtype.Timestamptz{
			Time:  refresh_token_exp_time,
			Valid: true,
		},
	})
	if err != nil {
		log.Println("here 7")
		return nil, err
	}

	err = tx.Commit(context.Background())
	if err != nil {
		return nil, err
	}

	res := AuthResponse{
		AccessToken:  access_token,
		RefreshToken: refresh_token,
	}

	return &res, nil

}
