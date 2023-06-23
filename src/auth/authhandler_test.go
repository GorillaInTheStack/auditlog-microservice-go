package auth

import (
	"auditlog/config"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/dgrijalva/jwt-go"
	qt "github.com/frankban/quicktest"
)

func GenerateTestToken() string {
	token := jwt.New(jwt.SigningMethodHS256)

	// Set the claims
	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix() // Token expiry time

	// Generate the token string
	tokenString, err := token.SignedString(config.SecretKey)
	if err != nil {
		log.Fatalf("Auth generateTokenTest: Authentication token generation error err: %v", err)
		return ""
	}

	return strings.Replace(tokenString, "Bearer", "", -1)
}

func TestAuthHandler(t *testing.T) {

	// Create a test HTTP server
	testServer := httptest.NewServer(AuthHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})))
	defer testServer.Close()

	t.Run("Valid token", func(t *testing.T) {

		c := qt.New(t)

		req, err := http.NewRequest(http.MethodGet, testServer.URL, nil)
		c.Assert(err, qt.IsNil, qt.Commentf("Auth Test: Error generating request %v", err))

		validToken := GenerateTestToken()
		log.Printf("Auth Test: generated valid token %v", validToken)

		req.Header.Set("Authorization", validToken)

		resp, err := http.DefaultClient.Do(req)
		c.Assert(err, qt.IsNil, qt.Commentf("Auth Test: Error sending request %v", err))
		defer resp.Body.Close()

		// Check the response status code
		c.Assert(resp.StatusCode, qt.Equals, http.StatusOK,
			qt.Commentf("Auth Test: Expected status code %d, got %d", http.StatusOK, resp.StatusCode))
	})

	t.Run("Missing token", func(t *testing.T) {

		c := qt.New(t)

		req, err := http.NewRequest(http.MethodGet, testServer.URL, nil)
		c.Assert(err, qt.IsNil, qt.Commentf("Auth Test: Error generating request %v", err))

		req.Header.Set("Authorization", "")

		resp, err := http.DefaultClient.Do(req)
		c.Assert(err, qt.IsNil, qt.Commentf("Auth Test: Error sending request %v", err))
		defer resp.Body.Close()

		// Check the response status code
		c.Assert(resp.StatusCode, qt.Equals, http.StatusUnauthorized,
			qt.Commentf("Auth Test: Expected status code %d, got %d", http.StatusUnauthorized, resp.StatusCode))
	})

	t.Run("Invalid token", func(t *testing.T) {

		c := qt.New(t)

		req, err := http.NewRequest(http.MethodGet, testServer.URL, nil)
		c.Assert(err, qt.IsNil, qt.Commentf("Auth Test: Error generating request %v", err))
		req.Header.Set("Authorization", "invalid-token")

		// Send the request
		resp, err := http.DefaultClient.Do(req)
		c.Assert(err, qt.IsNil, qt.Commentf("Auth Test: Error sending request %v", err))
		defer resp.Body.Close()

		// Verify the response status code
		if resp.StatusCode != http.StatusUnauthorized {
			t.Errorf("Auth Test: Expected status code %d, got %d", http.StatusUnauthorized, resp.StatusCode)
		}
	})
}
