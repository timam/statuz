package main

import (
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/statuzproj/statuz/cmd/watcher"
	"github.com/statuzproj/statuz/utils/healthz"
	"log"
	"net/http"
)

func main() {
	err := watcher.Start()
	if err != nil {
		log.Fatalf("Error starting watcher: %v", err)
	}

	http.HandleFunc("/healthz", healthz.HealthCheck)
	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":8080", nil)
}
