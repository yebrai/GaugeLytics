package postgres

import (
	"database/sql"
	"errors"
	"gaugelytics-backend/services/user-service/internal/domain"
	"gaugelytics-backend/services/user-service/internal/ports"

	_ "github.com/lib/pq"
)

type PostgresUserRepository struct {
	db *sql.DB
}

func NewPostgresUserRepository(db *sql.DB) ports.UserRepository {
	return &PostgresUserRepository{db: db}
}

func (r *PostgresUserRepository) Create(user domain.User) error {
	query := `
        INSERT INTO users (id, name, email, password, created_at)
        VALUES ($1, $2, $3, $4, $5)
    `
	_, err := r.db.Exec(query, user.ID, user.Name, user.Email, user.Password, user.CreatedAt)
	return err
}

func (r *PostgresUserRepository) FindByEmail(email string) (*domain.User, error) {
	query := `
        SELECT id, name, email, password, created_at
        FROM users
        WHERE email = $1
    `

	var user domain.User
	err := r.db.QueryRow(query, email).Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("user not found")
		}
		return nil, err
	}
	return &user, nil
}

// Otros métodos para actualizar y eliminar usuarios...
