package handlers

import (
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"

	"auditlog/config"
)

// generateToken generates a new JWT token
func GenerateTokenHandler(w http.ResponseWriter, r *http.Request) {
	token := jwt.New(jwt.SigningMethodHS256)

	// Set the claims
	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix() // Token expiry time
	// Add additional claims if needed

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
