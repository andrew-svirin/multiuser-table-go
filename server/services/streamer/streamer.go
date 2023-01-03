package streamer

// Streamer - streaming data into streams.

import (
	"log"
	"net/http"
	"os"
)

// StreamFromFile - Read file and stream it. Log error on fatal error.
func StreamFromFile(w http.ResponseWriter, fileName string) {
	bs, err := os.ReadFile(fileName)

	if err != nil {
		log.Fatal("No file exist: ", err)
	}

	_, err = w.Write(bs)

	if err != nil {
		log.Fatal("Response not written", err)
	}
}
