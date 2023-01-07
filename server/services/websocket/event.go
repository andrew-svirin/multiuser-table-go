package websocket

import (
	"encoding/json"
	"log"
)

const AuthorizeOp = "authorize"
const AuthorizedOp = "authorized"
const UserAuthorizedOp = "user/authorized"
const CellEditOp = "cell/edit"
const CellEditedOp = "cell/edited"
const UserCellEditedOp = "user/cell/edited"

// Event structure for exchange by websocket.
type Event struct {
	Op   string
	Data map[string]interface{}
}

// DecodeMessage - DecodeMessage event from message
func DecodeMessage(message []byte) *Event {
	var event Event

	err := json.Unmarshal(message, &event)

	if err != nil {
		log.Println("Message was not decoded", err)
	}

	return &event
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
func NewEvent(op string, data map[string]interface{}) *Event {
	return &Event{Op: op, Data: data}
}
