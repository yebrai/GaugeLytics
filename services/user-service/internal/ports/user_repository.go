package ports

import "gaugelytics-backend/services/user-service/internal/domain"

// UserRepository define las operaciones que se pueden realizar con los usuarios.
type UserRepository interface {
	Create(user domain.User) error
	FindByEmail(email string) (*domain.User, error)
}
