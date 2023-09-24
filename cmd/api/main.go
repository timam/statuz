package main

import (
	"log"
	"net/http"

	"github.com/timam/statuz/cmd/watcher"
	"github.com/timam/statuz/internal/healthz"
)

func main() {
	err := watcher.Start()
	if err != nil {
		log.Fatalf("Error starting watcher: %v", err)
	}

	http.HandleFunc("/healthz", healthz.HealthCheck)
	http.ListenAndServe(":8080", nil)
}
