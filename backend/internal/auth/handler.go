package auth

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/AlexanderWangY/swoppr-backend/db/sqlc"
	"github.com/AlexanderWangY/swoppr-backend/internal/utils"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgtype"
)

type Handler struct {
	service *UserService
}

func NewHandler(service *UserService) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) Register(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Extract the body JSON
	var register_req RegisterRequest
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&register_req)
	if err != nil {
		h.HandleError(w, "Wrong request format, check your json.", http.StatusBadRequest)
		return
	}

	hashed, err := utils.HashPassword(register_req.Password)
	if err != nil {
		h.HandleError(w, "Something went wrong.", http.StatusInternalServerError)
		return
	}

	tokens, err := h.service.RegisterUserAndJWT(sqlc.CreateUserParams{
		Email:           register_req.Email,
		PasswordHash:    hashed,
		FirstName:       pgtype.Text{Valid: false},
		LastName:        pgtype.Text{Valid: false},
		IsEmailVerified: false,
	})
	if err != nil {
		// Check for postgres errors
		if pgErr, ok := err.(*pgconn.PgError); ok && pgErr.Code == "23505" {
			h.HandleError(w, "That account already exists.", http.StatusConflict)
			return
		}

		log.Println(err)
		h.HandleError(w, "Something went wrong creating user.", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(tokens)
	if err != nil {
		h.HandleError(w, "Something went wrong sending response.", http.StatusInternalServerError)
	}
}

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Extract the body JSON
	var register_req RegisterRequest
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&register_req)
	if err != nil {
		h.HandleError(w, "Wrong request format, check your json.", http.StatusBadRequest)
		return
	}
}

func (h *Handler) GetOwnUser(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("This is a protected route you are accessing! This means you have a JWT!"))
	if err != nil {
		h.HandleError(w, "Something went wrong responding.", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *Handler) HandleError(w http.ResponseWriter, message string, status int) {
	errResponse := ErrorResponse{
		Error: message,
	}
	w.WriteHeader(status)
	err := json.NewEncoder(w).Encode(errResponse)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}
