package server

import (
	"github.com/andrew-svirin/multiuser-table-go/server/services/router"
	"github.com/andrew-svirin/multiuser-table-go/server/services/websocket"
	"net/http"
)

// WsServer - custom ws server structure.
type WsServer struct {
	Server
	upgrader       *websocket.Upgrader
	connectionPool *websocket.ConnectionPool
}

func (s *WsServer) CountConnections() int {
	return s.connectionPool.Count()
}

func NewWsServer(port int, router *router.WsRouter) *WsServer {
	u := websocket.NewUpgrader()
	cp := websocket.NewConnectionPool()

	// Ignore checking of client origin.
	co := func(r *http.Request) bool { return true }
	u.CheckOriginFunc(co)

	router.BeforeRouteCall = func(w http.ResponseWriter, r *http.Request) (*websocket.ConnectionPool, int) {
		c := u.UpgradeConnection(w, r)
		id := cp.Push(c)

		return cp, id
	}

	s := &WsServer{
		Server:         *NewServer(port, router),
		upgrader:       u,
		connectionPool: cp,
	}

	return s
}
