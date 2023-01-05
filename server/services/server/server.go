package server

import (
	"context"
	"net/http"
	"strconv"
)

// Server Server - custom server structure.
type Server struct {
	http.Server
}

// Serve - start listen and serve for incoming requests.
func (s *Server) Serve() {
	err := s.Server.ListenAndServe()

	if err != nil && err != http.ErrServerClosed {
		panic(err)
	}
}

// Shutdown - stop listen and serve for incoming requests.
func (s *Server) Shutdown() {
	err := s.Server.Shutdown(context.TODO())

	if err != nil {
		panic(err)
	}
}

// NewServer - instantiate new server.
func NewServer(port int, handler http.Handler) *Server {
	return &Server{
		http.Server{
			Addr:    ":" + strconv.Itoa(port),
			Handler: handler,
		},
	}
}
