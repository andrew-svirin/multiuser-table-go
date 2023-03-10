package server

import (
	"github.com/andrew-svirin/multiuser-table-go/server/services/router"
	"github.com/andrew-svirin/multiuser-table-go/server/services/websocket"
	"net/http"
)

// WsServer - custom ws server structure.
type WsServer struct {
	*Server
	upgrader       *websocket.Upgrader
	connectionPool *websocket.ConnectionPool
}

func (s *WsServer) CountConnections() int {
	return s.connectionPool.Count()
}

// NewWsServer - Instantiate new ws server.
func NewWsServer(pt string, ro *router.WsRouter) *WsServer {
	u := websocket.NewUpgrader()
	cp := websocket.NewConnectionPool()

	// Ignore checking of client origin.
	co := func(r *http.Request) bool { return true }
	u.CheckOriginFunc(co)

	ro.BeforeRouteCall = func(w http.ResponseWriter, r *http.Request) *websocket.Bus {
		c := u.UpgradeConnection(w, r)
		cId := cp.Push(c)

		return websocket.NewRequest(cp, cId)
	}

	return &WsServer{
		Server:         NewServer(pt, ro),
		upgrader:       u,
		connectionPool: cp,
	}
}
