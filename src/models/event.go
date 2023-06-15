package models

import (
	"time"
)

// Event represents the audit log event structure
type Event struct {
	EventID               string                 `bson:"EventID" json:"EventID"`                             // Event Unique ID generated by the audit log
	Timestamp             time.Time              `bson:"Timestamp" json:"Timestamp"`                         // Timestamp in nanoseconds of the moment the audit log received the event
	SourceEventID         string                 `bson:"SourceEventID" json:"SourceEventID"`                 // Event ID provided by the source service
	SourceTimestamp       time.Time              `bson:"SourceTimestamp" json:"SourceTimestamp"`             // Timestamp in nanoseconds provided by the source service
	CorrelationID         string                 `bson:"CorrelationID" json:"CorrelationID"`                 // Event correlation ID in case other events are related
	SourceTimezone        string                 `bson:"SourceTimezone" json:"SourceTimezone"`               // Timezone of source service
	SourceServiceName     string                 `bson:"SourceServiceName" json:"SourceServiceName"`         // Name of source service
	SourceServiceLocation string                 `bson:"SourceServiceLocation" json:"SourceServiceLocation"` // Source service geographical information
	SourceIPAddress       string                 `bson:"SourceIPAddress" json:"SourceIPAddress"`             // Source service IP address
	EventTags             map[string]string      `bson:"EventTags" json:"EventTags"`                         // Event metadata
	EventDataHash         string                 `bson:"EventDataHash" json:"EventDataHash"`                 // Event data hash or checksum for integrity
	EventDataVersion      string                 `bson:"EventDataVersion" json:"EventDataVersion"`           // Event data version
	EventData             map[string]interface{} `bson:"EventData" json:"EventData"`                         // Variant event data depending on their type
}
