package controller

import (
	"github.com/andrew-svirin/multiuser-table-go/server/services/websocket"
)

// HandleWebsocket - handling websocket messages.
// Dispatch to everyone incoming message.
func HandleWebsocket(cp *websocket.ConnectionPool, id int) int {
	mt, message := cp.ReadConnMessage(id)

	if message != nil {
		cp.DispatchMessage(mt, message)
	} else {
		cp.CloseConn(id)
	}

	return mt
}
