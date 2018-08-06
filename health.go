package main

import (
	"log"
	"net/http"
)

func health(w http.ResponseWriter, req *http.Request) {
	if req.Method != "GET" {
		w.Header().Set("accept", "GET")
		http.Error(w, http.StatusText(405), 405)
		return
	}

	log.Println("[soildesc] Health check")
	w.WriteHeader(http.StatusOK)
}
