package main

import (
	"fmt"
	"net/http"
)

func webHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, Golang!")
}

func main() {
	startServer()
}

func startServer() {
	http.HandleFunc("/", webHandler)

	fmt.Println("Server will start, you can access http://localhost:80")

	err := http.ListenAndServe(":80", nil)
	if err != nil {
		fmt.Printf("Failed to start server: %v", err)
	}
}