package runtime

import (
	"github.com/andrew-svirin/multiuser-table-go/server/services/router"
	"github.com/andrew-svirin/multiuser-table-go/server/services/server"
)

// initHttpServer - init HTTP server.
func (r *Runtime) initHttpServer() {
	ro := router.NewHttpRouter()
	r.config.HttpServer.RegisterRoutesCall(ro)
	r.httpServer = server.NewHttpServer(r.config.HttpServer.Port, ro)
}

// initWsServer - init WebSocket server.
func (r *Runtime) initWsServer() {
	ro := router.NewWsRouter()
	r.config.WsServer.RegisterRoutesCall(ro)
	r.wsServer = server.NewWsServer(r.config.WsServer.Port, ro)
}

// StartServers - initialize process of starting
// all servers.
func (r *Runtime) StartServers() {
	amount := 2 // amount of servers (http and socket)
	r.wg.Add(amount)
}

// ServeHttpServer - serve HTTP server.
func (r *Runtime) ServeHttpServer() {
	r.wg.Done()
	r.httpServer.Serve()
}

// ServeWsServer - serve Socket server.
func (r *Runtime) ServeWsServer() {
	r.wg.Done()
	r.wsServer.Serve()
}

// WaitServersStarted - Wait to run all serving servers.
// Should be put in the end of start servers process.
func (r *Runtime) WaitServersStarted() {
	r.wg.Wait()
}

// stopServers - Shutdown servers.
func (r *Runtime) stopServers() {
	r.httpServer.Shutdown()
	r.wsServer.Shutdown()
}
