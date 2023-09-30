package main

import (
	"github.com/statuzproj/statuz/utils/healthz"
	"log"
	"net/http"
)

func main() {
	log.Println("Work in Progress!!")
	http.HandleFunc("/healthz", healthz.HealthCheck)
	http.ListenAndServe(":8081", nil)
}
