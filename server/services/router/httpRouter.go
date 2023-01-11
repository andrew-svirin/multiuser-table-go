package router

// In this file registered all routes for http server.

import (
	"net/http"
)

// httpHandler - handler structure.
type httpHandler func(http.ResponseWriter, *http.Request)

// HttpRouter - custom router structure.
type HttpRouter struct {
	http.ServeMux
}

// AddRoute - add route handle for common page.
func (hr *HttpRouter) AddRoute(pattern string, h httpHandler) {
	hr.HandleFunc(pattern, h)
}

// AddIndexRoute - add route handle for index page.
func (hr *HttpRouter) AddIndexRoute(pattern string, h httpHandler) {
	indexRoute := func(w http.ResponseWriter, r *http.Request) {
		// "/" matches everything in index subtree,
		// thus we need to check that path is index page.
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}

		h(w, r)
	}

	hr.HandleFunc(pattern, indexRoute)
}

// NewHttpRouter - instantiate new http router.
func NewHttpRouter() *HttpRouter {
	return &HttpRouter{}
}
