package runtime

import "github.com/andrew-svirin/multiuser-table-go/server/services/router"

// Config - config for runtime.
type Config struct {
	HttpServer HttpServerConfig
	WsServer   WsServerConfig
}

// HttpServerConfig - config for http server.
type HttpServerConfig struct {
	Port               string
	RegisterRoutesCall func(*router.HttpRouter)
}

// WsServerConfig - config for websocket server.
type WsServerConfig struct {
	Port               string
	RegisterRoutesCall func(*router.WsRouter)
}
