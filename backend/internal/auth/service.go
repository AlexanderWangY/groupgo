package auth

import (
	"context"

	"github.com/AlexanderWangY/swoppr-backend/db"
	"github.com/AlexanderWangY/swoppr-backend/db/sqlc"
	"github.com/google/uuid"
)

type UserService struct {
	db *db.Database
}

func NewUserService(database *db.Database) *UserService {
	return &UserService{
		db: database,
	}
}

func (s *UserService) GetUser(uuid uuid.UUID) (sqlc.AuthUser, error) {
	var user sqlc.AuthUser
	var err error
	user, err = s.db.Query.GetUser(context.Background(), uuid)
	return user, err
}
