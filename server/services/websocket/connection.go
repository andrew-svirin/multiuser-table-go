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

// GoingAwayMessage - happens when connection was aborted.
const GoingAwayMessage = -1

// Push - add new connection into the pool.
func (cp *ConnectionPool) Push(c *websocket.Conn) int {
	cp.counter++

	cp.connections[cp.counter] = c

	log.Printf("Established Connection `%d`", cp.counter)

	return cp.counter
}

// take - takes specified connection from the pool.
func (cp *ConnectionPool) take(id int) *websocket.Conn {
	return cp.connections[id]
}

// ReadConnMessage - read message from connection in the pool.
func (cp *ConnectionPool) ReadConnMessage(id int) (int, []byte) {
	c := cp.take(id)

	if c == nil {
		return websocket.TextMessage, nil
	}

	mt, message, err := c.ReadMessage()

	if err != nil {
		log.Printf("Message can not be read by ID `%d`: %s", id, err)

		return mt, nil
	}

	return mt, message
}

// CloseConn - close specified connection in the pool.
func (cp *ConnectionPool) CloseConn(id int) {
	c := cp.take(id)

	c.Close()

	delete(cp.connections, id)

	log.Printf("Closed Connection `%d`", id)
}

// DispatchMessage - dispatch same message in the pool.
func (cp *ConnectionPool) DispatchMessage(messageType int, data []byte) {
	var err error

	for id, c := range cp.connections {
		err = c.WriteMessage(messageType, data)

		if err != nil {
			log.Printf("Message can not be written to ID `%d`: %s", id, err)
		}
	}
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
