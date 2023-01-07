package websocket

import (
	"github.com/gorilla/websocket"
	"log"
)

// ConnectionPool - pool of client websocket connections.
type ConnectionPool struct {
	counter     int
	connections map[int]*websocket.Conn
}

// GoingAwayMessage - uses when connection was aborted.
const GoingAwayMessage = -1

// TextMessage - uses when message is correct.
const TextMessage = 1

// Push - add new connection into the pool.
func (cp *ConnectionPool) Push(c *websocket.Conn) int {
	cp.counter++

	cp.connections[cp.counter] = c

	return cp.counter
}

// Get - get one connection from the pool.
func (cp *ConnectionPool) Get(id int) *websocket.Conn {
	return cp.connections[id]
}

// Delete - delete one connection in the pool.
func (cp *ConnectionPool) Delete(id int) {
	c := cp.Get(id)

	c.Close()

	delete(cp.connections, id)
}

// GetAll - get all connections from the pool.
func (cp *ConnectionPool) GetAll() *ConnectionCollection {
	return NewConnectionCollection(cp.connections)
}

// Count - count connections in pool.
func (cp *ConnectionPool) Count() int {
	return len(cp.connections)
}

// NewConnectionPool - instantiate new connection pool.
func NewConnectionPool() *ConnectionPool {
	return &ConnectionPool{
		counter:     0,
		connections: map[int]*websocket.Conn{},
	}
}

// ConnectionCollection - collection of client websocket connections.
// Represent interface of one client websocket.
type ConnectionCollection struct {
	connections map[int]*websocket.Conn
}

// WriteMessage - write same message to each connection.
func (cc *ConnectionCollection) WriteMessage(messageType int, data []byte) {
	var err error

	for id, c := range cc.connections {
		err = c.WriteMessage(messageType, data)

		if err != nil {
			log.Printf("Message can not be written to ID `%d`: %s", id, err)
		}
	}
}

// NewConnectionCollection - instantiate new connection collection.
func NewConnectionCollection(c map[int]*websocket.Conn) *ConnectionCollection {
	return &ConnectionCollection{
		connections: c,
	}
}
