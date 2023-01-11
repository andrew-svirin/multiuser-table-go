package runtime

import (
	"bufio"
	"fmt"
	"github.com/andrew-svirin/multiuser-table-go/server/services/system"
	"os"
	"strings"
)

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
		// Read command.
		text, _ := reader.ReadString('\n')

		// convert CRLF to LF
		text = strings.Replace(text, "\n", "", -1)

		switch text {
		case "help":
		case "h":
			fmt.Println("Allowed commands:\n" +
				"count (c)	- count connections\n" +
				"usage (u)	- usage of system\n" +
				"exit (e) 	- to exit from program")
			break
		case "count":
		case "c":
			fmt.Println("Connections:", r.wsServer.CountConnections())
			break
		case "usage":
		case "u":
			fmt.Printf("Usage:\nMemory = %vKb	CPU = %.2f%%\n",
				system.MemoryUsageKb(), system.CPUUsagePercents())
			break
		case "exit":
		case "e":
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
