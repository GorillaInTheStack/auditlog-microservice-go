package server_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"auditlog/config"
	"auditlog/server"

	qt "github.com/frankban/quicktest"
)

func TestStart(t *testing.T) {

	c := qt.New(t)

	// Set up the testing environment
	config.TestingEnabled = true
	config.Address = "127.0.0.1:7878"

	// Create a test HTTP server
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Handle the request
		server.Start()
	}))
	defer ts.Close()

	//test case: Test the "/test/auth" endpoint
	resp, err := http.Get(ts.URL + "/test/auth")
	c.Assert(err, qt.IsNil, qt.Commentf("Server Test: Error sending request %v", err))
	defer resp.Body.Close()

	// Assert the response status code
	c.Assert(resp.StatusCode, qt.Equals, http.StatusOK,
		qt.Commentf("Server Test: Expected status code %d, got %d", http.StatusOK, resp.StatusCode))

}
