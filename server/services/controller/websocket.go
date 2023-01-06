package controller

import (
	"github.com/andrew-svirin/multiuser-table-go/server/services/websocket"
)

// HandleWebsocket - handling websocket messages.
// Dispatch to everyone incoming message.
func HandleWebsocket(cp *websocket.ConnectionPool, id int) int {
	mt, message := cp.ReadConnMessage(id)

	if message != nil {
		ie := websocket.DecodeMessage(message)
		oe := websocket.NewEvent()

		handleEvent(ie, oe, cp, id)

		cp.DispatchMessage(mt, websocket.EncodeMessage(ie))
		cp.WriteConnMessage(id, mt, websocket.EncodeMessage(oe))
	} else {
		cp.CloseConn(id)
	}

	return mt
}

// handleEvent - route event.
// ie - incoming event
// oe - outgoing event
func handleEvent(ie *websocket.Event, oe *websocket.Event, cp *websocket.ConnectionPool, id int) {
	switch ie.Op {
	case websocket.AuthorizeOp:
		authorize(ie, oe, cp, id)
		break
	}
}

// Make event about authorized user.
func authorize(_ *websocket.Event, oe *websocket.Event, _ *websocket.ConnectionPool, id int) {
	oe.Op = websocket.AuthorizedOp
	oe.Data = map[string]interface{}{
		"type": "user",
		"id":   id,
	}
}
