package main

import (
	"github.com/andrew-svirin/multiuser-table-go/server/services/runtime"
	"log"
)

// main - Entrypoint.
func main() {
	r := new(runtime.Runtime)
	r.Init()

	r.StartServers()
	go func() {
		log.Println("HTTP Server started")

		r.ServeHttpServer()
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
