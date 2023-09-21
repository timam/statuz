package main

import (
	"github.com/timam/statuz/internal/healthz"
	"net/http"
)

func main() {
	http.HandleFunc("/healthz", healthz.HealthCheck)
	http.ListenAndServe(":8080", nil)
}
