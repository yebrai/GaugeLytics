package main

import (
	"database/sql"
	"fmt"
	"gaugelytics-backend/services/user-service/internal/application"
	"gaugelytics-backend/services/user-service/internal/config"
	"gaugelytics-backend/services/user-service/internal/infrastructure/db/postgres"
	"gaugelytics-backend/services/user-service/internal/infrastructure/http"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Could not load config: %v", err)
	}

	db, err := sql.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName))
	if err != nil {
		log.Fatalf("Could not connect to the database: %v", err)
	}
	defer db.Close()

	repo := postgres.NewPostgresUserRepository(db)
	service := application.NewUserService(repo)
	handler := http.NewUserHandler(service)

	http.HandleFunc("/users", handler.CreateUser)

	log.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
