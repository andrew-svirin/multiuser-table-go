package index

import (
	"../../services/streamer"
	"net/http"
)

// Handle - handling index route by reading file and streaming it.
func Handle(w http.ResponseWriter, _ *http.Request) {

	fileName := "./controllers/index/files/index.html"

	streamer.StreamFromFile(w, fileName)
}
