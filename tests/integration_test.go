package tests

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/your-org/my-go-app/internal/handlers"
)

func TestHealthEndpoint(t *testing.T) {
	req, err := http.NewRequest("GET", "/health", nil)
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handlers.HealthHandler)
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	var body map[string]interface{}
	err = json.Unmarshal(rr.Body.Bytes(), &body)
	assert.NoError(t, err)
	assert.Equal(t, "healthy", body["status"])
}

func TestReadyEndpoint(t *testing.T) {
	req, err := http.NewRequest("GET", "/ready", nil)
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handlers.ReadyHandler)
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	var body map[string]interface{}
	err = json.Unmarshal(rr.Body.Bytes(), &body)
	assert.NoError(t, err)
	assert.Equal(t, "ready", body["status"])
}

func TestHomeEndpoint(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handlers.HomeHandler("1.0.0"))
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	var body map[string]interface{}
	err = json.Unmarshal(rr.Body.Bytes(), &body)
	assert.NoError(t, err)
	assert.Contains(t, body["message"], "Golang EKS App")
	assert.Equal(t, "1.0.0", body["version"])
}

func TestInfoEndpoint(t *testing.T) {
	req, err := http.NewRequest("GET", "/info", nil)
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handlers.InfoHandler("1.0.0"))
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	var body map[string]interface{}
	err = json.Unmarshal(rr.Body.Bytes(), &body)
	assert.NoError(t, err)
	assert.Equal(t, "my-go-app", body["app_name"])
}
