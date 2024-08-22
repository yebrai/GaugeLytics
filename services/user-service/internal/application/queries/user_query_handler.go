package queries

import (
	"gaugelytics-backend/services/user-service/internal/domain"
	"gaugelytics-backend/services/user-service/internal/ports"
)

type UserQueryHandler struct {
	repo ports.UserRepository
}

func NewUserQueryHandler(repo ports.UserRepository) *UserQueryHandler {
	return &UserQueryHandler{repo: repo}
}

func (h *UserQueryHandler) GetUserByEmail(email string) (*domain.User, error) {
	return h.repo.FindByEmail(email)
}
