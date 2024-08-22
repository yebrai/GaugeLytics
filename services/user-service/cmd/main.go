package main

import (
	"database/sql"
	"log"
	"net/http"

	"gaugelytics-backend/services/user-service/internal/application"
	"gaugelytics-backend/services/user-service/internal/infrastructure/db/postgres"
	_ "github.com/lib/pq" // PostgreSQL driver
)

func main() {
	// Configuración de la base de datos
	connStr := "user=postgres dbname=user_service sslmode=disable" // Ajusta según tu configuración
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("error opening database connection: %v", err)
	}
	defer db.Close()

	userRepo := postgres.NewPostgresUserRepository(db)
	createUserHandler := application.NewCreateUserCommandHandler(userRepo)

	mux := http.NewServeMux()

	mux.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			var cmd application.CreateUserCommand
			err := createUserHandler.Handle(cmd)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			w.WriteHeader(http.StatusCreated)
		default:
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		}
	})

	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatalf("error starting server: %v", err)
	}
}
