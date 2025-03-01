package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/AlexanderWangY/swoppr-backend/db"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Something went wrong loading the .env")
	}

	db.ConnectDB()
	defer db.CloseDB()

	r := chi.NewRouter()

	r.Use(middleware.Logger)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome!"))
	})

	r.Get("/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("This is the hello endpoints!"))
	})

	if err := http.ListenAndServe(":3000", r); err != nil {
		fmt.Println("Error starting server:", err)
	} else {
		fmt.Println("Listening on port 3000")
	}
}
