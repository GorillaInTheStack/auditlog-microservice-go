package services

import (
	"auditlog/models"
	"auditlog/storage"
	"log"
)

type LocalService struct {
}

func (l *LocalService) SaveEvent(event models.Event) error {
	// Implementation for saving event in local storage
	err := storage.InsertEvent(event)
	if err != nil {
		log.Println("Service: Failed to save event in local storage:", err)
	} else {
		log.Println("Service: Event saved in local storage:", event.EventID)
	}
	return err
}

func (l *LocalService) GetEvents(key string, value interface{}) ([]models.Event, error) {
	// Implementation for retrieving event from local storage
	events, err := storage.GetEventByKeyValue(key, value)
	if err != nil {
		log.Println("Service: Failed to retrieve events from local storage:", err)
	}
	return events, err
}
