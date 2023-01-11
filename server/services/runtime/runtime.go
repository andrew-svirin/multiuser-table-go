package runtime

// Runtime - is interface to handle main project
// sub routines. Commands should be run in certain order.
// Consist of servers managements and CMD shell.

import (
	"github.com/andrew-svirin/multiuser-table-go/server/services/server"
	"sync"
)

// Runtime - object that aggregate dependents.
type Runtime struct {
	httpServer *server.HttpServer
	wsServer   *server.WsServer
	wg         *sync.WaitGroup
	config     *Config
}

// Init - init runtime dependencies.
// Run it at first initialization.
func (r *Runtime) Init() {
	r.wg = &sync.WaitGroup{}
	r.initHttpServer()
	r.initWsServer()
}

// stop - Correct closing all runtime dependencies.
func (r *Runtime) stop() {
	r.stopServers()
}

// NewRuntime - instantiate new runtime.
func NewRuntime(config *Config) *Runtime {
	return &Runtime{
		config: config,
	}
}
