package services

import (
	"auditlog/config"
	"auditlog/models"
)

var eventService EventService

func init() {

	if config.IsClustered {
		eventService = &RemoteService{}
	} else {
		eventService = &LocalService{}
	}
}

// This function saves an event using an event service and returns an error if there is one.
func SaveEvent(event models.Event) error {
	err := eventService.SaveEvent(event)
	return err
}

// The function retrieves events based on a given key-value pair.
func GetEventsByKeyValue(key string, value interface{}) ([]models.Event, error) {
	event, err := eventService.GetEvents(key, value)

	if err != nil {
		return nil, err
	}

	// Event retrieved successfully
	return event, nil
}
