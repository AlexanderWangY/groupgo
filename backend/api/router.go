package api

import (
	"log"
	"net/http"

	"github.com/AlexanderWangY/swoppr-backend/db"
	"github.com/AlexanderWangY/swoppr-backend/internal/auth"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func SetupRouter(database *db.Database) *chi.Mux {

	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)

	authService := auth.NewUserService(database)
	authHandler := auth.NewHandler(authService)

	r.Route("/api/v1", func(r chi.Router) {

		r.Route("/auth", func(r chi.Router) {
			r.Post("/register", authHandler.Register)
		})
	})

	return r
}

func StartServer(database *db.Database) {
	r := SetupRouter(database)

	log.Println("Starting server on port 8080")
	http.ListenAndServe(":8080", r)
}
