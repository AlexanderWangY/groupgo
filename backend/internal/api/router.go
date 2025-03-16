package api

import (
	"context"
	"log"
	"net/http"

	"github.com/AlexanderWangY/swoppr-backend/internal/db"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func SetupRouter(database *db.Database) *chi.Mux {

	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)

	r.Route("/api/v1", func(r chi.Router) {

		r.Route("/auth", func(r chi.Router) {

			// Public /auth routes
			r.Group(func(r chi.Router) {
			})

			// Private /auth routes
			r.Group(func(r chi.Router) {
			})
		})
	})

	return r
}

func StartServer(ctx context.Context, database *db.Database) {
	r := SetupRouter(database)

	log.Println("Starting server on port 8080")
	http.ListenAndServe(":8080", r)
}
