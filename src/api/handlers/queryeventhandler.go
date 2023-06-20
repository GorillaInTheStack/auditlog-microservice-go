package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"auditlog/models"
	"auditlog/services"
)

var (
	GetEventsByKeyValueService = services.GetEventsByKeyValue
)

// This function retrieves events based on query parameters, stores unique events in a map, converts
// the map to a slice, encodes the events as JSON, and sends the JSON response.
func QueryEventHandler(w http.ResponseWriter, r *http.Request) {
	// Parse the query parameters
	queryParams := r.URL.Query()

	if len(queryParams) == 0 {
		http.Error(w, "No filter given", http.StatusBadRequest)
		log.Println("Handler: No query Params given")
		return
	}

	// Create a map to store unique events
	totalevents := make(map[string]models.Event)

	for key, values := range queryParams {
		// If there are multiple values for a key,  we'll pick the first value
		value := values[0]

		// Call the service function to retrieve events
		events, err := GetEventsByKeyValueService(key, value)

		if err != nil {
			// Handle error and send appropriate response
			http.Error(w, "Error retrieving events", http.StatusInternalServerError)
			log.Printf("Handler: Error retrieving events: %s\n", err)
			return
		}

		// Add events to a map, using the event ID as the key
		for _, event := range events {
			totalevents[event.EventID] = event
		}
	}

	// Convert the map to a slice
	var uniqueEvents []models.Event
	for _, event := range totalevents {
		uniqueEvents = append(uniqueEvents, event)
	}

	// Encode events as JSON
	jsonData, err := json.Marshal(uniqueEvents)

	if err != nil {
		// Handle error and send appropriate response
		http.Error(w, "Error encoding events", http.StatusInternalServerError)
		log.Printf("Handler: Error encoding events: %s\n", err)
		return
	}

	// Set the response headers
	w.Header().Set("Content-Type", "application/json")

	log.Printf("Handler: Returing filtered events: %v", uniqueEvents)

	// Write the JSON response
	w.Write(jsonData)

}
