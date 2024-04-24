package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	log.Println("Starting server...")

	server := http.NewServeMux()

	server.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "<h1>Hello, World!</h1>")
	})
	
	log.Println("Started server on port :8080")
	if err := http.ListenAndServe("0.0.0.0:8080", server); err != nil {
		log.Fatal(err)
	}
}
