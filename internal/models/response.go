package models

import "time"

type HealthResponse struct {
	Status    string    `json:"status"`
	Timestamp time.Time `json:"timestamp"`
}

type HomeResponse struct {
	Message   string    `json:"message"`
	Version   string    `json:"version"`
	Hostname  string    `json:"hostname"`
	Timestamp time.Time `json:"timestamp"`
}

type InfoResponse struct {
	AppName   string    `json:"app_name"`
	Version   string    `json:"version"`
	Hostname  string    `json:"hostname"`
	GoVersion string    `json:"go_version"`
	Timestamp time.Time `json:"timestamp"`
}

type ErrorResponse struct {
	Error   string `json:"error"`
	Message string `json:"message"`
}
