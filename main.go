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

func health(w http.ResponseWriter, req *http.Request) {
	if req.Method != "GET" {
		w.Header().Set("Allow", "POST")
		http.Error(w, http.StatusText(405), 405)
		return
	}

	w.WriteHeader(http.StatusOK)
}
