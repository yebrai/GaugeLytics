package commands

import (
	"gaugelytics-backend/services/user-service/internal/domain"
	"gaugelytics-backend/services/user-service/internal/ports"
	"time"
)

type UserCommandHandler struct {
	repo ports.UserRepository
}

func NewUserCommandHandler(repo ports.UserRepository) *UserCommandHandler {
	return &UserCommandHandler{repo: repo}
}

func (h *UserCommandHandler) CreateUser(name, email, password string) error {
	user := domain.User{
		ID:        generateID(),
		Name:      name,
		Email:     email,
		Password:  hashPassword(password),
		CreatedAt: time.Now(),
	}
	return h.repo.Create(user)
}

func (h *UserCommandHandler) UpdateUser(user domain.User) error {
	return h.repo.Update(user)
}

func (h *UserCommandHandler) DeleteUser(id string) error {
	return h.repo.Delete(id)
}

// Utility functions
func generateID() string {
	return "some-unique-id" // Replace with actual ID generation logic
}

func hashPassword(password string) string {
	return "hashed-password" // Replace with actual password hashing logic
}
