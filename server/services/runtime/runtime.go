package runtime

// Runtime - is interface to handle main project
// sub routines. Commands should be run in certain order.
// Consist of servers managements and CMD shell.

import (
	"bufio"
	"fmt"
	"github.com/andrew-svirin/multiuser-table-go/server/services/router"
	"github.com/andrew-svirin/multiuser-table-go/server/services/server"
	"os"
	"strings"
	"sync"
)

// Runtime - object that aggregate dependents.
type Runtime struct {
	httpServer *server.HttpServer
	wg         *sync.WaitGroup
}

// Init - init runtime dependencies.
// Run it at first initialization.
func (r *Runtime) Init() {
	r.wg = new(sync.WaitGroup)
	r.initHttpServer()
}

// initHttpServer - init HTTP server.
func (r *Runtime) initHttpServer() {
	ro := router.ResolveHttpRouter()

	r.httpServer = new(server.HttpServer)

	r.httpServer.Init(8080, ro)
}

// StartServers - initialize process of starting
// all servers.
func (r *Runtime) StartServers() {
	r.wg.Add(1)
}

// ServeHttpServer - serve HTTP server.
func (r *Runtime) ServeHttpServer() {
	r.wg.Done()
	r.httpServer.Serve()
}

// WaitServersStarted - Wait to run all serving servers.
// Should be put in the end of start servers process.
func (r *Runtime) WaitServersStarted() {
	r.wg.Wait()
}

// StartCmd - initialize process of starting
// cmd .
func (r *Runtime) StartCmd() {
	r.wg.Add(1)
}

// ServeCmd - handles cmd shell commands
// to manage runtime dependencies.
func (r *Runtime) ServeCmd() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("CMD Shell\n---------------------")

	for {
		fmt.Print("-> ")

		// Read command.
		text, _ := reader.ReadString('\n')

		// convert CRLF to LF
		text = strings.Replace(text, "\n", "", -1)

		switch text {
		case "help":
			fmt.Println("Allowed commands:\n" +
				"exit - to exit from program")
			break
		case "exit":
			fmt.Println("Exiting...")
			r.stop()
			r.wg.Done()
			return
		}
	}
}

// WaitCmdExit - should be put on the end of
// serve cmd process.
func (r *Runtime) WaitCmdExit() {
	r.wg.Wait()
}

// stop - Correct closing all runtime dependencies.
func (r *Runtime) stop() {
	r.httpServer.Shutdown()
}
