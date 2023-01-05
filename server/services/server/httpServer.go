package server

import (
	"github.com/andrew-svirin/multiuser-table-go/server/services/router"
)

// HttpServer - custom http server structure.
type HttpServer struct {
	*Server
}

// NewHttpServer - initiate new http server.
func NewHttpServer(port int, router *router.HttpRouter) *HttpServer {
	return &HttpServer{
		Server: NewServer(port, router),
	}
}
