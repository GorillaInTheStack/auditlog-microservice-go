package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	qt "github.com/frankban/quicktest"
)

func TestGenerateTokenHandler(t *testing.T) {
	tests := []struct {
		name         string
		expectedCode int
	}{
		{
			name:         "Valid token generation",
			expectedCode: http.StatusOK,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			c := qt.New(t)

			req, err := http.NewRequest(http.MethodGet, "/generatetoken", nil)
			c.Assert(err, qt.IsNil, qt.Commentf("Handler Test: Error generating request %v", err))

			rr := httptest.NewRecorder()

			GenerateTokenHandler(rr, req)

			// Check the response status code
			c.Assert(rr.Code, qt.Equals, test.expectedCode,
				qt.Commentf("Handler Test: Expected status code %d, got %d", test.expectedCode, rr.Code))

		})
	}

}
