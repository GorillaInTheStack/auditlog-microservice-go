package storage

import (
	"auditlog/models"
	"log"
	"reflect"
)

var eventsMap map[string]models.Event

// This function initializes a map for local storage.
func init() {
	// Using a map for local storage for demo purposes ONLY.
	eventsMap = make(map[string]models.Event)
}

// The function inserts an event into a map and logs the event ID.
func InsertEvent(event models.Event) error {
	eventsMap[event.EventID] = event
	log.Printf("Storage: Event inserted locally: %s", event.EventID)

	return nil
}

// The function filters events by a given key-value pair and returns the filtered events.
func GetEventByKeyValue(key string, value interface{}) ([]models.Event, error) {
	var filteredEvents []models.Event

	for _, event := range eventsMap {

		eventValue := reflect.ValueOf(event)
		fieldValue := eventValue.FieldByName(key)

		if event.EventData[key] == value || (fieldValue.IsValid() && fieldValue.String() == value) {
			filteredEvents = append(filteredEvents, event)
		}

	}

	log.Printf("Storage: Filtered events by key-value locally: %s=%s", key, value)
	return filteredEvents, nil
}

// The function retrieves an event by its ID and returns it along with a boolean indicating if it was
// found.
func GetEventByID(eventID string) (models.Event, bool) {
	event, found := eventsMap[eventID]
	if found {
		log.Printf("Storage: Event found locally: %s", event.EventID)
	} else {
		log.Printf("Storage: Event not found locally: %s", eventID)
	}
	return event, found
}

func DeleteEvent() {
	// TODO: too busy, implement later
}
func UpdateEvent() {
	// TODO: too busy, implement later
}
