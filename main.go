package main

import (
	"./controllers/index"
	"log"
	"net/http"
)

// main - Enter point.
func main() {
	startServer()
}

// startServer - Register router and start listen for incoming requests.
func startServer() {
	http.HandleFunc("/", index.Handle)

	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
