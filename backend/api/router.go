package api

import (
	"net/http"

	"github.com/AlexanderWangY/swoppr-backend/db"
	"github.com/AlexanderWangY/swoppr-backend/db/sqlc"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func SetupRouter() http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)

	queries := sqlc.New(db.DB)
}
