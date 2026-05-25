package services

import (
	"os"
	"runtime"
	"time"

	"github.com/your-org/my-go-app/internal/models"
)

func GetHealthStatus() models.HealthResponse {
	return models.HealthResponse{
		Status:    "healthy",
		Timestamp: time.Now(),
	}
}

func GetReadyStatus() models.HealthResponse {
	return models.HealthResponse{
		Status:    "ready",
		Timestamp: time.Now(),
	}
}

func GetHomeResponse(version string) models.HomeResponse {
	hostname, _ := os.Hostname()
	return models.HomeResponse{
		Message:   "Hello from Golang EKS App deployed via Harness!",
		Version:   version,
		Hostname:  hostname,
		Timestamp: time.Now(),
	}
}

func GetInfoResponse(version string) models.InfoResponse {
	hostname, _ := os.Hostname()
	return models.InfoResponse{
		AppName:   "my-go-app",
		Version:   version,
		Hostname:  hostname,
		GoVersion: runtime.Version(),
		Timestamp: time.Now(),
	}
}
