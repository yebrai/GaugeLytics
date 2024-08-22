package unit

import (
	"encoding/json"
	"gaugelytics-backend/services/health-service/internal/application"
	"github.com/gorilla/mux"
	"net/http"
	"net/http/httptest"
	"testing"
)

type HealthCheckResponse struct {
	Status string `json:"status"`
}

func TestHealthCheckHandler(t *testing.T) {
	r := mux.NewRouter()
	r.HandleFunc("/health", application.HealthCheckHandler).Methods("GET")

	req, err := http.NewRequest("GET", "/health", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	var response HealthCheckResponse
	err = json.NewDecoder(rr.Body).Decode(&response)
	if err != nil {
		t.Fatalf("failed to decode response body: %v", err)
	}

	expected := HealthCheckResponse{Status: "healthy"}

	if response != expected {
		t.Errorf("handler returned unexpected body: got %+v want %+v", response, expected)
	}
}
