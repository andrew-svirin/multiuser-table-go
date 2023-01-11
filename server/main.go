package main

import (
	"github.com/andrew-svirin/multiuser-table-go/server/app"
)

var a *app.App

// init - Invoking before main func.
func init() {
	a = app.NewApp()
	a.Init()
}

// main - Entrypoint.
func main() {
	a.Open()
	defer a.Close()

	a.Run()
}
