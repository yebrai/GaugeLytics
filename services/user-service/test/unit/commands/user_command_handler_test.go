package commands

import (
	"gaugelytics-backend/services/user-service/internal/application/commands"
	"gaugelytics-backend/services/user-service/internal/domain"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

// Mock del repositorio de usuarios
type MockUserRepository struct {
	mock.Mock
}

func (m *MockUserRepository) Create(user domain.User) error {
	args := m.Called(user)
	return args.Error(0)
}

func TestCreateUserCommandHandler(t *testing.T) {
	mockRepo := new(MockUserRepository)
	handler := commands.NewUserCommandHandler(mockRepo)

	user := domain.User{ID: "1", Name: "Test User", Email: "test@example.com"}

	mockRepo.On("Create", user).Return(nil)

	err := handler.HandleCreateUserCommand(user)
	assert.NoError(t, err)

	mockRepo.AssertExpectations(t)
}
