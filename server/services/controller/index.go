package controller

// Controller for index page.

import (
	"github.com/andrew-svirin/multiuser-table-go/server/services/filesystem"
	"net/http"
)

// HandleIndex - handling index route by reading file and streaming it.
func HandleIndex(w http.ResponseWriter, r *http.Request) {
	fileName := filesystem.ResolveResourcePath("/index.html")

	http.ServeFile(w, r, fileName)
}
