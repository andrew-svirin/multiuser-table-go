package main

import (
	"github.com/andrew-svirin/multiuser-table-go/server/services/runtime"
	"log"
)

var r *runtime.Runtime

// init - Invoking before main func.
func init() {
	r = new(runtime.Runtime)
	r.Init()
}

// main - Entrypoint.
func main() {
	r.StartServers()
	go func() {
		log.Println("HTTP Server started")

		r.ServeHttpServer()
	}()

	go func() {
		log.Println("WS Server started")

		r.ServeWsServer()
	}()

	r.WaitServersStarted()

	r.StartCmd()
	go func() {
		log.Println("Cmd Server started")

		r.ServeCmd()
	}()
	r.WaitCmdExit()

	log.Println("Exit")
}
