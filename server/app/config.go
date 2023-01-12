package app

import (
	"github.com/andrew-svirin/multiuser-table-go/server/controllers"
	"github.com/andrew-svirin/multiuser-table-go/server/services/database"
	"github.com/andrew-svirin/multiuser-table-go/server/services/router"
	"github.com/andrew-svirin/multiuser-table-go/server/services/runtime"
)

// Config - App config.
type Config struct {
	Runtime runtime.Config
	Db      database.Config
}

// NewConfig - instantiate new config.
func NewConfig() *Config {
	return &Config{
		Runtime: runtime.Config{
			HttpServer: runtime.HttpServerConfig{
				Port: "8080",
				RegisterRoutesCall: func(ro *router.HttpRouter) {
					ro.AddIndexRoute("/", controllers.HandleIndex)
					ro.AddRoute("/static/", controllers.HandleStatic)
				},
			},
			WsServer: runtime.WsServerConfig{
				Port: "3456",
				RegisterRoutesCall: func(ro *router.WsRouter) {
					ro.AddIndexRoute("/", controllers.HandleWebsocket)
					ro.AddEventRoute("authorize", controllers.HandleAuthorize)
					ro.AddEventRoute("cell/save", controllers.HandleCellSave)
					ro.AddEventRoute("cell/load/all", controllers.HandleCellLoadAll)
				},
			},
		},
		Db: database.Config{
			Driver: "mysql",
			Dsn:    "root:password@tcp(host.docker.internal:3306)/multiuser_table?parseTime=true",
		},
	}
}
