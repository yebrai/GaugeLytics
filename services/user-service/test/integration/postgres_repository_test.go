package integration

import (
	"database/sql"
	"gaugelytics-backend/services/user-service/internal/domain"
	"gaugelytics-backend/services/user-service/internal/infrastructure/db/postgres"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateUserIntegration(t *testing.T) {
	db, err := sql.Open("postgres", "user=postgres dbname=testdb sslmode=disable")
	if err != nil {
		t.Fatal(err)
	}
	repo := postgres.NewPostgresUserRepository(db)

	user := domain.User{ID: "1", Name: "Test User", Email: "test@example.com"}
	err = repo.Create(user)
	assert.NoError(t, err)

	// Verifica la creación del usuario consultando la base de datos
	createdUser, err := repo.FindByEmail(user.Email)
	assert.NoError(t, err)
	assert.Equal(t, user, *createdUser)
}
