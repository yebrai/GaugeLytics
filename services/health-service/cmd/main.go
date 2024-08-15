package main

import (
	"gaugelytics-backend/services/health-service/internal/service"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/health", service.HealthCheckHandler)
	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Error starting server: %v\n", err)
	}
}
