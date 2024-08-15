package service

import (
	"encoding/json"
	"gaugelytics-backend/services/health-service/internal/domain"
	"net/http"
)

// HealthCheckHandler handles the /health endpoint
func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	response := domain.HealthCheck{Status: "healthy"}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
