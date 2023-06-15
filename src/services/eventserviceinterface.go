package services

import "auditlog/models"

// The EventService interface defines methods for saving and retrieving events.
// @property {error} SaveEvent - SaveEvent is a method that belongs to an interface called
// EventService. It takes in a parameter of type models.Event and returns an error. The purpose of this
// method is to save an event to some kind of storage, such as a database.
// @property GetEvents - GetEvents is a method of the EventService interface that takes in a key and a
// value as parameters and returns a slice of Event objects and an error. It is used to retrieve events
// from a data source based on a specific key-value pair.
// This interface exists mainly to support both local and remote services.
type EventService interface {
	SaveEvent(event models.Event) error
	GetEvents(key string, value interface{}) ([]models.Event, error)
}
