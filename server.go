package main

import (
	"fmt"
	"io"
	"net/http"
)

func requestHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello, CodeSnippet.io!")
}

func main() {
	// register the request handler,
	// which accepts requests at the /api route
	http.HandleFunc("/api", requestHandler)

	// log some details to the console
	fmt.Println("Server listing on port 4500...")

	// start the web server process to begin handling requests
	http.ListenAndServe(":4500", nil)
}
