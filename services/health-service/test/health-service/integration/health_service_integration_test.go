// services/health-service/test/health-service/integration/health_service_integration_test.go

package service

import (
	"log"
	"net/http"
	"testing"
	"time"
)

func TestHealthCheckEndpoint(t *testing.T) {
	// Start the service in a goroutine
	go func() {
		err := http.ListenAndServe(":8080", nil)
		if err != nil {
			log.Fatalf("Error starting server: %v\n", err)
		}
	}()

	// Wait for the server to start
	time.Sleep(2 * time.Second)

	// Perform the HTTP request
	resp, err := http.Get("http://localhost:8080/health")
	if err != nil {
		t.Fatalf("Failed to make request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status code 200, got %v", resp.StatusCode)
	}

	// Optionally, check the response body
	// expected := `{"status":"healthy"}`
	// body, _ := io.ReadAll(resp.Body)
	// if string(body) != expected {
	//     t.Errorf("Expected response body %v, got %v", expected, string(body))
	// }
}
