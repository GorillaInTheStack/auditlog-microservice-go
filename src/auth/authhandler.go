package auth

import (
	"log"
	"net/http"

	"github.com/dgrijalva/jwt-go"

	"auditlog/config"
)

// AuthHandler is a middleware for authenticating requests
func AuthHandler(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		authHeader := r.Header.Get("Authorization")

		if authHeader == "" {
			log.Println("Auth: Received request with authentication token missing")
			http.Error(w, "Authentication token is missing", http.StatusUnauthorized)
			return
		}

		token, err := jwt.Parse(authHeader, func(token *jwt.Token) (interface{}, error) {
			return config.SecretKey, nil
		})
		if err != nil || !token.Valid {
			log.Println("Auth: Received invalid token")
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		next(w, r)
	}
}
