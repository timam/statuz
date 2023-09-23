package main

import (
	"github.com/timam/statuz/config"
	"github.com/timam/statuz/internal/healthz"
	"log"
	"net/http"
)

func main() {
	_, err := config.GetEnvVars()
	if err != nil {
		log.Fatalln(err)
	}

	log.Println("starting statuz watcher")

	http.HandleFunc("/healthz", healthz.HealthCheck)
	http.ListenAndServe(":8080", nil)
}
