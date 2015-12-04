package eventsource

import "time"

type Event struct {
	Id        string      `json:"id"`
	Stream    string      `json:"stream"`
	CreatedOn time.Time   `json:"created_on"`
	Body      interface{} `json:"body"`
}

type EventStore interface {
	Send(stream string, eventData interface{}) (string, error)
}
