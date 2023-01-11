package app

import (
	"github.com/andrew-svirin/multiuser-table-go/server/services/runtime"
	"log"
)

// initRuntime - init runtime.
func (a *App) initRuntime() {
	a.runtime = runtime.NewRuntime(&a.config.Runtime)

	a.runtime.Init()
}

// runRuntime - run runtime.
func (a *App) runRuntime() {
	a.runtime.StartServers()
	go func() {
		log.Println("HTTP Server started")

		a.runtime.ServeHttpServer()
	}()

	go func() {
		log.Println("WS Server started")

		a.runtime.ServeWsServer()
	}()

	a.runtime.WaitServersStarted()

	a.runtime.StartCmd()
	go func() {
		log.Println("Cmd Server started")

		a.runtime.ServeCmd()
	}()
	a.runtime.WaitCmdExit()
}
