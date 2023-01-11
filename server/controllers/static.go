package controllers

import (
	"github.com/andrew-svirin/multiuser-table-go/server/services/filesystem"
	"net/http"
)

// HandleStatic - handling static route by reading file and streaming it.
func HandleStatic(w http.ResponseWriter, r *http.Request) {
	dirName := filesystem.ResolveStaticPath("/")

	fs := http.FileServer(http.Dir(dirName))

	http.StripPrefix(filesystem.StaticDir+"/", fs).ServeHTTP(w, r)
}
