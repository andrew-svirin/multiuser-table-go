package controller

import (
	"github.com/andrew-svirin/multiuser-table-go/server/services/websocket"
)

// HandleWebsocket - handling websocket messages.
// Dispatch to everyone incoming message.
func HandleWebsocket(b *websocket.Bus) int {
	mt, ie := b.ConnectionReadEvent()

	if ie == nil {
		b.ConnectionDelete()
	} else {
		handleEvent(ie, b)
	}

	return mt
}

// handleEvent - route event.
func handleEvent(ie *websocket.Event, b *websocket.Bus) {
	switch ie.Op {
	case websocket.AuthorizeOp:
		authorize(ie, b)
		break
	case websocket.CellEditOp:
		cellEdit(ie, b)
		break
	}
}

// Handle event about authorized user.
func authorize(_ *websocket.Event, b *websocket.Bus) {
	oe := websocket.NewEvent(
		websocket.AuthorizedOp,
		map[string]interface{}{
			"id": b.ConnectionId(),
		},
	)
	b.ConnectionWriteEvent(websocket.TextMessage, oe)

	oea := websocket.NewEvent(
		websocket.UserAuthorizedOp,
		map[string]interface{}{
			"id": b.ConnectionId(),
		},
	)
	go b.ConnectionPoolWriteEvent(websocket.TextMessage, oea)
}

// Handle event about cell edit.
func cellEdit(ie *websocket.Event, b *websocket.Bus) {
	oe := websocket.NewEvent(
		websocket.CellEditedOp,
		map[string]interface{}{
			"name":  ie.Data["name"],
			"value": ie.Data["value"],
		},
	)
	b.ConnectionWriteEvent(websocket.TextMessage, oe)

	oea := websocket.NewEvent(
		websocket.UserCellEditedOp,
		map[string]interface{}{
			"user_id": b.ConnectionId(),
			"name":    ie.Data["name"],
			"value":   ie.Data["value"],
		},
	)
	go b.ConnectionPoolWriteEvent(websocket.TextMessage, oea)
}
