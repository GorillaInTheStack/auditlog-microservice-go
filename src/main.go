package main

import (
	"log"

	"auditlog/server"
)

/*
var (
	secretKey = []byte("test") //TODO: Should be retrived from env and kept safely.
	events    []Event
)

func filterEvents(eventType, userID string) []Event {
	var filteredEvents []Event
	//for _, event := range events {
		//if (eventType == "" || event.EventType == eventType) &&
		//	(userID == "" || event.UserID == userID) {
		//	filteredEvents = append(filteredEvents, event)
		//}
	//}
	return filteredEvents
}

*/

func main() {
	log.Println("Main: Server starting...")
	server.Start()
	log.Println("Main: Server shutdown")
}
