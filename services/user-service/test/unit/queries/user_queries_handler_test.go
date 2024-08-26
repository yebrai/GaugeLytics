package queries

import (
	"gaugelytics-backend/services/user-service/internal/application/queries"
	"gaugelytics-backend/services/user-service/internal/domain"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

// Mock del repositorio de usuarios
type MockUserRepository struct {
	mock.Mock
}

func (m *MockUserRepository) FindByEmail(email string) (*domain.User, error) {
	args := m.Called(email)
	return args.Get(0).(*domain.User), args.Error(1)
}

func TestFindUserQueryHandler(t *testing.T) {
	mockRepo := new(MockUserRepository)
	handler := queries.NewUserQueryHandler(mockRepo)

	email := "test@example.com"
	user := &domain.User{ID: "1", Name: "Test User", Email: email}

	mockRepo.On("FindByEmail", email).Return(user, nil)

	result, err := handler.HandleFindUserQuery(email)
	assert.NoError(t, err)
	assert.Equal(t, user, result)

	mockRepo.AssertExpectations(t)
}
