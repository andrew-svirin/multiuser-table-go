package websocket

import (
	"github.com/gorilla/websocket"
	"net/http"
)

// Upgrader - custom upgrader structure.
type Upgrader struct {
	websocket.Upgrader
}

// UpgradeConnection - establish new connection.
func (u *Upgrader) UpgradeConnection(w http.ResponseWriter, r *http.Request) *websocket.Conn {
	c, err := u.Upgrade(w, r, nil)

	if err != nil {
		panic(err)
	}

	return c
}

// CheckOriginFunc - uses for checking request origin.
// Can be used to identify client.
func (u *Upgrader) CheckOriginFunc(handleFunc func(r *http.Request) bool) {
	u.CheckOrigin = handleFunc
}

// NewUpgrader - instantiate new upgrader.
func NewUpgrader() *Upgrader {
	u := new(Upgrader)

	return u
}
