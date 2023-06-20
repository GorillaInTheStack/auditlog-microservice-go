package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"auditlog/models"
	"auditlog/services"

	"github.com/google/uuid"
)

var (
	EventSaverService = services.SaveEvent
)

// This function handles the submission of an event, decodes the event data, saves the event, and logs
// any errors or successful saves.
func SubmitEventHandler(w http.ResponseWriter, r *http.Request) {

	decoder := json.NewDecoder(r.Body)

	var event models.Event
	err := decoder.Decode(&event)

	if err != nil {
		http.Error(w, "Invalid event data", http.StatusBadRequest)
		log.Printf("Handler: Error decoding event data while submitting event: %s\n", err)
		return
	}
	defer r.Body.Close()

	event.EventID = uuid.New().String()
	event.Timestamp = time.Now()

	log.Printf("Handler: Received event from service eventID: %s\n", event.EventID)
	log.Printf("Handler: Decoded event: %v", event)

	err = EventSaverService(event)

	if err != nil {
		log.Printf("Handler: Error saving eventID: %s\n %s\n", event.EventID, err)
		w.WriteHeader(http.StatusBadRequest)
	} else {
		log.Printf("Handler: Event saved successfully eventID: %s\n", event.EventID)
		w.WriteHeader(http.StatusOK)
	}
}
