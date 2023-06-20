package handlers

import (
	"auditlog/models"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	qt "github.com/frankban/quicktest"
)

type QueryEventTest struct {
	name            string
	queryParams     map[string][]string
	expectedCode    int
	expectedEvents  []models.Event
	expectedHttpLog []string
}

func TestQueryEventHandler(t *testing.T) {

	tests := []QueryEventTest{
		{
			name: "Valid Query Parameters",
			queryParams: map[string][]string{
				"EventID":          {"event2"},
				"EventDataVersion": {"1"},
			},
			expectedCode: http.StatusOK,
			expectedEvents: []models.Event{
				{EventID: "1", SourceEventID: "event1", EventDataVersion: "1"},
				{EventID: "2", SourceEventID: "event2", SourceServiceLocation: "London"},
			},
			expectedHttpLog: []string{
				"event1",
				"event2",
			},
		},
		{
			name:         "Empty Query Parameters",
			queryParams:  map[string][]string{},
			expectedCode: http.StatusBadRequest,
			expectedEvents: []models.Event{
				{EventID: "1", SourceEventID: "event1", EventDataVersion: "1"},
				{EventID: "2", SourceEventID: "event2", SourceServiceLocation: "London"},
			},
			expectedHttpLog: []string{"No filter given"},
		},
		{
			name: "Get Event Service Error",
			queryParams: map[string][]string{
				"EventID":          {"event2"},
				"EventDataVersion": {"1"},
			},
			expectedCode:    http.StatusInternalServerError,
			expectedEvents:  []models.Event{},
			expectedHttpLog: []string{"Error retrieving events"},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			c := qt.New(t)

			// Create a test HTTP request with the given request body
			req, err := http.NewRequest(http.MethodGet, "/events/query", nil)

			c.Assert(err, qt.IsNil)

			q := req.URL.Query()
			for key, values := range test.queryParams {
				q[key] = values
			}
			req.URL.RawQuery = q.Encode()
			fmt.Println(req.URL.RawQuery)
			rr := httptest.NewRecorder()

			// Mock get event service execution
			GetEventsByKeyValueService = func(key string, value interface{}) ([]models.Event, error) {

				if len(test.expectedEvents) > 0 {
					return test.expectedEvents, nil
				} else {
					return nil, errors.New("Error retrieving events")
				}

			}

			QueryEventHandler(rr, req)

			// Check the response status code
			c.Assert(rr.Code, qt.Equals, test.expectedCode,
				qt.Commentf("Expected status code %d, got %d", test.expectedCode, rr.Code))

			// Check the logged message
			logOutput := rr.Body.String()

			var events []models.Event
			err = json.Unmarshal([]byte(logOutput), &events)

			var extractedHttpOutput string

			if err != nil {
				// no events returned
				extractedHttpOutput = logOutput
			} else {
				SourceEventIDs := []string{
					events[0].SourceEventID,
					events[1].SourceEventID,
				}

				extractedHttpOutput = strings.Join(SourceEventIDs, ", ")
			}

			c.Assert(extractedHttpOutput, qt.Contains, strings.Join(test.expectedHttpLog, ", "),
				qt.Commentf("Expected http message '%s' not found in http output: %s", strings.Join(test.expectedHttpLog, ", "), extractedHttpOutput))
		})
	}

}
