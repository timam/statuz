package main

import (
	"github.com/timam/statuz/internal/healthz"
	"log"
	"net/http"
)

func main() {
	log.Println("starting statuz watcher")
	http.HandleFunc("/healthz", healthz.HealthCheck)
	http.ListenAndServe(":8080", nil)
}
