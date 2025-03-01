package auth

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

type Handler struct {
	service *UserService
}

func NewHandler(service *UserService) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) GetUser(w http.ResponseWriter, r *http.Request) {
	user_id := chi.URLParam(r, "userId")
	w.Header().Set("Content-Type", "application/json")
	// Parse the UUID from the string
	user_uuid, err := uuid.Parse(user_id)
	if err != nil {
		h.HandleError(w, "Not a valid UUID", http.StatusBadRequest)
		return
	}

	user, err := h.service.GetUser(user_uuid)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			h.HandleError(w, "No rows found!", http.StatusNotFound)
		} else {
			h.HandleError(w, "Internal server error.", http.StatusInternalServerError)
		}
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}

func (h *Handler) HandleError(w http.ResponseWriter, message string, status int) {
	errResponse := ErrorResponse{
		Error: message,
	}
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(errResponse)
}
