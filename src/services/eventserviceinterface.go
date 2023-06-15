package services

import "auditlog/models"

type EventService interface {
	SaveEvent(event models.Event) error
	GetEvents(key string, value interface{}) ([]models.Event, error)
}
