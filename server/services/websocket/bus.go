package websocket

import (
	"github.com/gorilla/websocket"
	"log"
)

// Aggregated structure for ws request.

type Bus struct {
	connectionPool *ConnectionPool
	connectionId   int
}

// ConnectionPool - getter for connection pool.
func (r *Bus) ConnectionPool() *ConnectionPool {
	return r.connectionPool
}

// ConnectionId - getter for current connection id.
func (r *Bus) ConnectionId() int {
	return r.connectionId
}

// connectionReadMessage - read message from current connection.
func (r *Bus) connectionReadMessage() (int, []byte) {
	c := r.connectionPool.Get(r.connectionId)

	mt, message, err := c.ReadMessage()

	if err != nil {
		log.Printf("Message can not be read by ID `%d`: %s", r.connectionId, err)

		return mt, nil
	}

	return websocket.TextMessage, message
}

// ConnectionReadEvent - read event from current connection.
func (r *Bus) ConnectionReadEvent() (int, *Event) {
	mt, msg := r.connectionReadMessage()

	if mt != websocket.TextMessage {
		return mt, nil
	}

	e := DecodeMessage(msg)

	return mt, &e
}

// ConnectionDelete - delete current connection.
func (r *Bus) ConnectionDelete() {
	r.connectionPool.Delete(r.connectionId)
}

// connectionWriteMessage - write message into current connection.
func (r *Bus) connectionWriteMessage(mt int, message []byte) {
	err := r.connectionPool.Get(r.connectionId).WriteMessage(mt, message)

	if err != nil {
		log.Printf("Message can not be written to ID `%d`: %s", r.connectionId, err)
	}
}

// ConnectionWriteEvent - write event into current connection.
func (r *Bus) ConnectionWriteEvent(e *Event) {
	message := EncodeEvent(e)

	r.connectionWriteMessage(websocket.TextMessage, message)
}

// connectionPoolWriteMessage - write message into connection pool.
func (r *Bus) connectionPoolWriteMessage(mt int, message []byte) {
	r.connectionPool.GetAll().WriteMessage(mt, message)
}

// ConnectionPoolWriteEvent - write event into connection pool.
func (r *Bus) ConnectionPoolWriteEvent(e *Event) {
	message := EncodeEvent(e)

	r.connectionPoolWriteMessage(websocket.TextMessage, message)
}

// NewRequest - creator for new  request.
func NewRequest(cp *ConnectionPool, cId int) *Bus {
	return &Bus{
		connectionPool: cp,
		connectionId:   cId,
	}
}
