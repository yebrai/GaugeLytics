package http

import (
	"encoding/json"
	"gaugelytics-backend/services/user-service/internal/application/commands"
	"gaugelytics-backend/services/user-service/internal/application/queries"
	"net/http"
	"time"
)

type UserHandler struct {
	commandHandler *commands.UserCommandHandler
	queryHandler   *queries.UserQueryHandler
}

func NewUserHandler(commandHandler *commands.UserCommandHandler, queryHandler *queries.UserQueryHandler) *UserHandler {
	return &UserHandler{commandHandler: commandHandler, queryHandler: queryHandler}
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	err := h.commandHandler.CreateUser(req.Name, req.Email, req.Password)
	if err != nil {
		http.Error(w, "Failed to create user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	email := r.URL.Query().Get("email")
	user, err := h.queryHandler.GetUserByEmail(email)
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	response := struct {
		ID        string `json:"id"`
		Name      string `json:"name"`
		Email     string `json:"email"`
		CreatedAt string `json:"created_at"`
	}{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		CreatedAt: user.CreatedAt.Format(time.RFC3339),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// Otros métodos para actualizar y eliminar usuarios...
