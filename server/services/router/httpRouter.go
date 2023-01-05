package router

// In this file registered all routes for http server.

import (
	"net/http"
)

// HttpRouter - custom router structure.
type HttpRouter struct {
	http.ServeMux
}

// AddRoute - add route handle for common page.
func (r *HttpRouter) AddRoute(pattern string, handler func(http.ResponseWriter, *http.Request)) {
	r.HandleFunc(pattern, handler)
}

// AddIndexRoute - add route handle for index page.
func (r *HttpRouter) AddIndexRoute(pattern string, handler func(http.ResponseWriter, *http.Request)) {
	indexRoute := func(w http.ResponseWriter, r *http.Request) {
		// "/" matches everything in index subtree,
		// thus we need to check that path is index page.
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}

		handler(w, r)
	}

	r.HandleFunc(pattern, indexRoute)
}

// NewHttpRouter - instantiate new http router.
func NewHttpRouter() *HttpRouter {
	return new(HttpRouter)
}
