package main

import (
	"github.com/andrew-svirin/multiuser-table-go/server/services/router"
	"github.com/andrew-svirin/multiuser-table-go/server/services/server"
)

// main - Entrypoint.
func main() {
	startHttpServer()
}

// startHttpServer - starts http server needs for handle http requests.
func startHttpServer() {
	r := router.ResolveHttpRouter()

	server.RunHttpServer(r)
}
