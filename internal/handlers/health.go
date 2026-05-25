package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/your-org/my-go-app/internal/services"
)

func writeJSON(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

func HomeHandler(version string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		resp := services.GetHomeResponse(version)
		writeJSON(w, http.StatusOK, resp)
	}
}

func HealthHandler(w http.ResponseWriter, r *http.Request) {
	resp := services.GetHealthStatus()
	writeJSON(w, http.StatusOK, resp)
}

func ReadyHandler(w http.ResponseWriter, r *http.Request) {
	resp := services.GetReadyStatus()
	writeJSON(w, http.StatusOK, resp)
}

func InfoHandler(version string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		resp := services.GetInfoResponse(version)
		writeJSON(w, http.StatusOK, resp)
	}
}
