package index

// Controller for index page.

import (
	"github.com/andrew-svirin/multiuser-table-go/server/services/filesystem"
	"github.com/andrew-svirin/multiuser-table-go/server/services/streamer"
	"net/http"
)

// Handle - handling index route by reading file and streaming it.
func Handle(w http.ResponseWriter, _ *http.Request) {

	fileName := filesystem.ResolveStaticPath("/index.html")

	streamer.StreamFromFile(w, fileName)
}
