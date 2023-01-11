package router

import (
	"github.com/andrew-svirin/multiuser-table-go/server/services/websocket"
	"log"
)

// EventHandler - event handler struct.
type EventHandler func(*websocket.Event, *websocket.Bus)

// EventRouter - event handler map.
type EventRouter struct {
	handlers map[string]EventHandler
}

// SetRoute - set route to the map.
func (ro *EventRouter) SetRoute(op string, h EventHandler) {
	ro.handlers[op] = h
}

// getHandler - get route handler.
func (ro *EventRouter) getHandler(op string) EventHandler {
	return ro.handlers[op]
}

// HandleEvent - handle event by handler.
func (ro *EventRouter) HandleEvent(ie *websocket.Event, b *websocket.Bus) {
	h := ro.getHandler(ie.Op)

	if h == nil {
		log.Println("Unhandled event", ie)
	} else {
		h(ie, b)
	}
}

// NewEventRouter - initialization for event router.
func NewEventRouter() *EventRouter {
	return &EventRouter{
		handlers: map[string]EventHandler{},
	}
}
