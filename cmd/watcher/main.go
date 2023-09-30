package main

import (
	"errors"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/statuzproj/statuz/config"
	"github.com/statuzproj/statuz/utils/healthz"
	"log"
	"net/http"
)

func main() {
	err := start()
	if err != nil {
		log.Fatalf("Error starting watcher: %v", err)
	}

	http.HandleFunc("/healthz", healthz.HealthCheck)
	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":8080", nil)
}

func start() error {
	env, err := config.GetEnvVars()
	if err != nil {
		return err
	}

	if env.Type == "webpage" {
		err := watchPage(env.Endpoint, env.Interval)
		if err != nil {
			return err
		}
	} else if env.Type == "ip" {
		watchIp(env.Endpoint, env.Port, env.Interval)
	} else {
		return errors.New("Unsupported watch type: " + env.Type)
	}
	return nil
}
