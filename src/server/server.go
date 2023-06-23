package server

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"auditlog/api/handlers"
	"auditlog/auth"
	"auditlog/config"
)

var server *http.Server

// This function starts an HTTP server and sets up handlers for different endpoints.
func Start() {
	http.HandleFunc("/generatetoken", handlers.GenerateTokenHandler)
	http.HandleFunc("/events/submit", auth.AuthHandler(handlers.SubmitEventHandler))
	http.HandleFunc("/events/query", auth.AuthHandler(handlers.QueryEventHandler))

	if config.TestingEnabled {
		http.HandleFunc("/test/auth", auth.AuthHandler(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
		}))
	}

	server = &http.Server{Addr: config.Address}

	// Start the HTTP server in a separate goroutine
	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server: Server error: %s", err)
		}
	}()

	// Wait for an interrupt signal to gracefully shut down the server
	if config.TestingEnabled {
		// Sleep for 2 seconds
		time.Sleep(2 * time.Millisecond)
	} else {
		quit := make(chan os.Signal, 1)
		signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
		<-quit
	}

	ShutDown()
}

// ShutDown gracefully shuts down the server
func ShutDown() {
	if server == nil {
		return
	}

	log.Println("Server: Shutting down server...")

	// Create a context with a timeout to allow the server to shut down gracefully
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Shut down the server
	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Server: Server shutdown error: %s", err)
	}

	log.Println("Server: Server gracefully stopped")
}
