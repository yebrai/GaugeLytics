package main

import (
	"gaugelytics-backend/services/health-service/internal/application"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// Crear un nuevo enrutador
	router := mux.NewRouter()

	// Definir la ruta para el health check
	router.HandleFunc("/health", application.HealthCheckHandler).Methods("GET")

	log.Println("Starting server on :8080")
	// Usar el enrutador con ListenAndServe
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatalf("Error starting server: %v\n", err)
	}
}
