package router

import (
	"github.com/andrew-svirin/multiuser-table-go/server/services/websocket"
	"net/http"
)

type WsRouter struct {
	http.ServeMux
	BeforeRouteCall func(w http.ResponseWriter, r *http.Request) (*websocket.ConnectionPool, int)
}

// AddIndexRoute - add route handle for index page.
// Where `id` - is newly added connection for client.
func (wsr *WsRouter) AddIndexRoute(pattern string, handler func(cp *websocket.ConnectionPool, id int) int) {
	indexRoute := func(w http.ResponseWriter, r *http.Request) {
		// "/" matches everything in index subtree,
		// thus we need to check that path is index page.
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}

		if wsr.BeforeRouteCall == nil {
			panic("BeforeRouteCall is not declared")
		}

		cp, id := wsr.BeforeRouteCall(w, r)

		// Listening for messages from client.
		for {
			mt := handler(cp, id)

			if mt == websocket.GoingAwayMessage {
				break
			}
		}
	}

	wsr.HandleFunc(pattern, indexRoute)
}

// NewWsRouter - instantiate new web socket router.
func NewWsRouter() *WsRouter {
	return new(WsRouter)
}
