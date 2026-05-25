package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/mux"
	"github.com/your-org/my-go-app/internal/handlers"
	"github.com/your-org/my-go-app/pkg/utils"
)

func main() {
	logger := utils.NewLogger()
	defer func() {
		_ = logger.Sync()
	}()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	appVersion := os.Getenv("APP_VERSION")
	if appVersion == "" {
		appVersion = "1.0.0"
	}

	router := mux.NewRouter()

	// Routes
	router.HandleFunc("/", handlers.HomeHandler(appVersion)).Methods("GET")
	router.HandleFunc("/health", handlers.HealthHandler).Methods("GET")
	router.HandleFunc("/ready", handlers.ReadyHandler).Methods("GET")
	router.HandleFunc("/info", handlers.InfoHandler(appVersion)).Methods("GET")

	srv := &http.Server{
		Addr:         ":" + port,
		Handler:      router,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	// Start server in goroutine
	go func() {
		fmt.Printf("Server starting on port %s | version %s\n", port, appVersion)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Sugar().Fatalf("Server failed: %v", err)
		}
	}()

	// Graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	fmt.Println("Shutting down server...")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		logger.Sugar().Fatalf("Server forced to shutdown: %v", err)
	}

	fmt.Println("Server exited")
}