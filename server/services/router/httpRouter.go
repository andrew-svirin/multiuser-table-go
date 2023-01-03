package router

// In this file registered all routes for http server.

import (
	"github.com/andrew-svirin/multiuser-table-go/server/controllers/index"
	"net/http"
)

// Router - custom router structure.
type Router struct {
	http.ServeMux
}

// AddRoute - add route handle for common page.
func (r *Router) AddRoute(pattern string, handler func(http.ResponseWriter, *http.Request)) {
	r.HandleFunc(pattern, handler)
}

// AddIndexRoute - add route handle for index page.
func (r *Router) AddIndexRoute(pattern string, handler func(http.ResponseWriter, *http.Request)) {
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

// ResolveHttpRouter - resolve router with defined routes for http server.
func ResolveHttpRouter() *Router {
	r := new(Router)

	r.AddIndexRoute("/", index.Handle)

	return r
}
