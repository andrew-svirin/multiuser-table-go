package app

// Application contains global and runtime instance
//and manage whole application by them.

import (
	"github.com/andrew-svirin/multiuser-table-go/server/services/runtime"
)

// App - application struct.
type App struct {
	config  *Config
	runtime *runtime.Runtime
}

// Init - initialize hook for app.
func (a *App) Init() {
	a.initDb()
	a.initRuntime()
}

// Open - open hook for app.
func (a *App) Open() {
	a.openDb()
}

// Close - close hook for app.
func (a *App) Close() {
	a.closeDb()
}

// Run - run hook for app.
func (a *App) Run() {
	a.runRuntime()
}

// NewApp - instantiate new application.
func NewApp() *App {
	return &App{
		config: NewConfig(),
	}
}
