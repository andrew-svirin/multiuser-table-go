package server

// Http server - listener on http port for requests
// and responding in answer.

import (
	"github.com/andrew-svirin/multiuser-table-go/server/services/router"
	"log"
	"net/http"
	"strconv"
)

// Server - custom server structure.
type Server struct {
	httpServer *http.Server
}

// Serve - Register router and Start listen incoming requests
func (s *Server) Serve(port int, handler http.Handler) {
	err := http.ListenAndServe(":"+strconv.Itoa(port), handler)

	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

// RunHttpServer - Serve http server.
func RunHttpServer(r *router.Router) {
	srv := new(Server)

	srv.Serve(8080, r)
}
