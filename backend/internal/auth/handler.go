package auth

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/AlexanderWangY/swoppr-backend/db/sqlc"
	"github.com/AlexanderWangY/swoppr-backend/internal/utils"
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
		log.Println(err)
		h.HandleError(w, "Something went wrong creating user.", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(tokens)

}

func (h *Handler) HandleError(w http.ResponseWriter, message string, status int) {
	errResponse := ErrorResponse{
		Error: message,
	}
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(errResponse)
}
