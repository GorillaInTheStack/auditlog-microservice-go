package handlers

import (
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"

	"auditlog/config"
)

// This function generates a JWT token with a specified expiry time and additional claims if needed.
func GenerateTokenHandler(w http.ResponseWriter, r *http.Request) {
	token := jwt.New(jwt.SigningMethodHS256)

	// Set the claims
	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix() // Token expiry time

	// Generate the token string
	tokenString, err := token.SignedString(config.SecretKey)
	if err != nil {
		http.Error(w, "Authentication token generation error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(tokenString))
}
