package router

import (
	"github.com/andrew-svirin/multiuser-table-go/server/services/websocket"
	"net/http"
)

type WsHandler func(*websocket.Bus, *EventRouter) int

// WsRouter - websocket router struct.
type WsRouter struct {
	http.ServeMux
	BeforeRouteCall func(http.ResponseWriter, *http.Request) *websocket.Bus
	eventRouter     *EventRouter
}

// AddIndexRoute - add route handle for index page.
// Where `id` - is newly added connection for client.
func (wsr *WsRouter) AddIndexRoute(pattern string, handler WsHandler) {
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

		b := wsr.BeforeRouteCall(w, r)

		// Listening for messages from client.
		for {
			mt := handler(b, wsr.eventRouter)

			if mt == websocket.GoingAwayMessage {
				break
			}
		}
	}

	wsr.HandleFunc(pattern, indexRoute)
}

// AddEventRoute - add event route handler.
func (wsr *WsRouter) AddEventRoute(op string, handler EventHandler) {
	wsr.eventRouter.SetRoute(op, handler)
}

// NewWsRouter - instantiate new web socket router.
func NewWsRouter() *WsRouter {
	return &WsRouter{
		eventRouter: NewEventRouter(),
	}
}
