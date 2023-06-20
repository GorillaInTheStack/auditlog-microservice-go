package handlers

import (
	"auditlog/models"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	qt "github.com/frankban/quicktest"
)

type SubmitEventTest struct {
	name            string
	requestBody     string
	expectedCode    int
	expectedHttpLog string
}

func TestSubmitEventHandler(t *testing.T) {
	tests := []SubmitEventTest{
		{
			name:            "Valid Event",
			requestBody:     `{"SourceEventID": "123456", "SourceTimestamp": "2023-06-11T08:30:00Z", "EventData": {"name": "John Doe"}}`,
			expectedCode:    http.StatusOK,
			expectedHttpLog: "",
		},
		{
			name:            "Invalid Event Data",
			requestBody:     `{"SourceEventID": "123456", "SourceTimestamp": "2023-06-11T08:30:00Z", "EventData": "invalid"}`,
			expectedCode:    http.StatusBadRequest,
			expectedHttpLog: "Invalid event data",
		},
		{
			name:            "Empty Request Body",
			requestBody:     "",
			expectedCode:    http.StatusBadRequest,
			expectedHttpLog: "Invalid event data",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			c := qt.New(t)

			// Create a test HTTP request with the given request body
			req, err := http.NewRequest(http.MethodPost, "/events/submit", strings.NewReader(test.requestBody))

			c.Assert(err, qt.IsNil)

			rr := httptest.NewRecorder()

			SubmitEventHandler(rr, req)

			// Check the response status code
			c.Assert(rr.Code, qt.Equals, test.expectedCode,
				qt.Commentf("Expected status code %d, got %d", test.expectedCode, rr.Code))

			// Check the logged message
			logOutput := rr.Body.String()
			c.Assert(logOutput, qt.Contains, test.expectedHttpLog,
				qt.Commentf("Expected http message '%s' not found in http output: %s", test.expectedHttpLog, logOutput))
		})
	}

	t.Run("Error while calling saveEvent service", func(t *testing.T) {

		c := qt.New(t)

		SavingErrorTest := SubmitEventTest{
			name:            "Event Saving Error",
			requestBody:     `{"SourceEventID": "123456", "SourceTimestamp": "2023-06-11T08:30:00Z", "EventData": {"name": "John Doe"}}`,
			expectedCode:    http.StatusBadRequest,
			expectedHttpLog: "",
		}

		// Create a test HTTP request with the given request body
		req, err := http.NewRequest(http.MethodPost, "/events/submit",
			strings.NewReader(SavingErrorTest.requestBody))
		rr := httptest.NewRecorder()

		c.Assert(err, qt.IsNil)

		// Mock event save service error
		EventSaverService = func(event models.Event) error {
			return errors.New("Service Mock: Error saving the event!")
		}

		SubmitEventHandler(rr, req)

		// Check the response status code
		c.Assert(rr.Code, qt.Equals, SavingErrorTest.expectedCode,
			qt.Commentf("Expected status code %d, got %d", SavingErrorTest.expectedCode, rr.Code))

	})
}
