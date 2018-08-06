package main

import (
	"log"
	"net/http"
)

func main() {

	log.Printf("Starting HTTP server on port 8000.\n")
	log.Printf("Press CTRL+C to stop.")
	http.HandleFunc("/health", health)
	http.HandleFunc("/describe", describe)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
