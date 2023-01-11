package controllers

import (
	"github.com/andrew-svirin/multiuser-table-go/server/services/websocket"
)

// HandleAuthorize - handle event about authorized user.
func HandleAuthorize(_ *websocket.Event, b *websocket.Bus) {
	oe := websocket.NewEvent(
		"authorized",
		websocket.EventData{
			"id": b.ConnectionId(),
		},
	)
	b.ConnectionWriteEvent(oe)

	oea := websocket.NewEvent(
		"user/authorized",
		websocket.EventData{
			"id": b.ConnectionId(),
		},
	)
	go b.ConnectionPoolWriteEvent(oea)
}
