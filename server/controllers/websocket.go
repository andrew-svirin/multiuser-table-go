package controllers

import (
	"github.com/andrew-svirin/multiuser-table-go/server/services/router"
	"github.com/andrew-svirin/multiuser-table-go/server/services/websocket"
)

// HandleWebsocket - handling websocket messages.
// Dispatch to everyone incoming message.
func HandleWebsocket(b *websocket.Bus, ro *router.EventRouter) int {
	mt, ie := b.ConnectionReadEvent()

	if ie == nil {
		handleDisconnect(b)
	} else {
		ro.HandleEvent(ie, b)
	}

	return mt
}

// handleDisconnect - handle user disconnect.
func handleDisconnect(b *websocket.Bus) {
	b.ConnectionDelete()

	oea := websocket.NewEvent(
		"user/disconnected",
		websocket.EventData{
			"id": b.ConnectionId(),
		},
	)
	go b.ConnectionPoolWriteEvent(oea)
}
