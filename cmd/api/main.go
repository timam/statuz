package main

import (
	"github.com/timam/statuz/cmd/watcher"
	"github.com/timam/statuz/internal/healthz"
	"net/http"
)

func main() {
	watcher.Start()

	http.HandleFunc("/healthz", healthz.HealthCheck)
	http.ListenAndServe(":8080", nil)
}
