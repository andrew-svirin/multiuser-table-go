package websocket

import (
	"encoding/json"
	"log"
)

type EventOperation string
type EventData map[string]interface{}

// Event structure for exchange by websocket.
type Event struct {
	Op   string
	Data EventData
}

// DecodeMessage - DecodeMessage event from message
func DecodeMessage(message []byte) Event {
	var event Event

	err := json.Unmarshal(message, &event)

	if err != nil {
		log.Println("Message was not decoded", err)
	}

	return event
}

// EncodeEvent - encode event to message
func EncodeEvent(e *Event) []byte {
	b, err := json.Marshal(e)

	if err != nil {
		log.Println(err)
	}

	return b
}

// NewEvent - creates new event.
func NewEvent(op string, d EventData) *Event {
	return &Event{Op: op, Data: d}
}
