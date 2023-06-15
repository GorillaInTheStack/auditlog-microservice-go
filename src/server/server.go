package server

import (
	"log"
	"net/http"

	"auditlog/auth"
	"auditlog/config"
	"auditlog/handlers"
)

// This function starts an HTTP server and sets up handlers for different endpoints.
func Start() {
	http.HandleFunc("/generatetoken", handlers.GenerateTokenHandler)
	http.HandleFunc("/events/submit", auth.AuthHandler(handlers.SubmitEventHandler))
	http.HandleFunc("/events/query", auth.AuthHandler(handlers.QueryEventHandler))

	// Start the HTTP server
	log.Printf("Server: listening on %s", config.Address)
	err := http.ListenAndServe(config.Address, nil)
	if err != nil {
		log.Fatalf("Server: encountered an error: %s", err)
	}
}
