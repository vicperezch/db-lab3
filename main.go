package main

import (
	"fmt"
	"net/http"
)

func messageHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	fmt.Fprintln(w, "pong")
}

func main() {
	http.HandleFunc("/ping", messageHandler)

	fmt.Println("Server starting on http://localhost:8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Server failed:", err)
	}
}
