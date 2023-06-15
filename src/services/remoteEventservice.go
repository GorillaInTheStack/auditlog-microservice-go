package services

import (
	"auditlog/models"
	"auditlog/storage"
	"log"
)

type RemoteService struct {
}

func (r *RemoteService) SaveEvent(event models.Event) error {
	// Implementation for saving event in remote storage
	err := storage.InsertDoc(event)
	if err != nil {
		log.Println("Service: Failed to save event in remote storage:", err)
	} else {
		log.Println("Service: Event saved in remote storage:", event.EventID)
	}
	return err
}

func (r *RemoteService) GetEvents(key string, value interface{}) ([]models.Event, error) {
	filter := map[string]interface{}{key: value}
	events, err := storage.FindDoc(filter)
	if err != nil {
		log.Println("Service: Failed to retrieve events from remote storage:", err)
		return nil, err
	}

	/*
		filteredEvents := make([]models.Event, len(events))
		for i, event := range events {
			if typedEvent, ok := event.(models.Event); ok {
				filteredEvents[i] = typedEvent
			} else {
				// Handle the case where the event is not of type models.Event
				// You can choose to skip or log the invalid events.
				log.Println("Service: Encountered event from remote storage that is not of type models.Event:", event)
			}
		}
	*/
	log.Printf("Service: Retrieved %d events from remote storage for key-value: %s=%v", len(events), key, value)
	return events, nil
}
