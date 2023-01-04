package server

// Http server - listener on http port for requests
// and responding in answer.

import (
	"context"
	"net/http"
	"strconv"
)

// HttpServer - custom server structure.
type HttpServer struct {
	server *http.Server
}

// Init - Register router and listening port.
func (s *HttpServer) Init(port int, handler http.Handler) {
	s.server = &http.Server{
		Addr:    ":" + strconv.Itoa(port),
		Handler: handler,
	}
}

// Serve - Start listen incoming requests.
func (s *HttpServer) Serve() {
	err := s.server.ListenAndServe()

	if err != nil && err.Error() != "http: Server closed" {
		panic(err)
	}
}

// Shutdown - Stop listening incoming requests.
func (s *HttpServer) Shutdown() {
	err := s.server.Shutdown(context.TODO())

	if err != nil {
		panic(err)
	}
}
