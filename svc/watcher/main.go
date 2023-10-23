package main

import (
	"github.com/statuzproj/statuz/utils/healthz"
	"log"
	"net/http"
)

//
//func main() {
//	var wg sync.WaitGroup
//
//	// Start the HTTP server in a goroutine
//	wg.Add(1)
//	go func() {
//		defer wg.Done()
//
//		http.HandleFunc("/healthz", healthz.HealthCheck)
//		http.Handle("/metrics", promhttp.Handler())
//		err := http.ListenAndServe(":8080", nil)
//		if err != nil {
//			log.Fatalf("HTTP server error: %v", err)
//		}
//	}()
//
//	// Start the watcher in another goroutine
//	wg.Add(1)
//	go func() {
//		defer wg.Done()
//
//		err := start()
//		if err != nil {
//			log.Fatalf("Error starting watcher: %v", err)
//		}
//	}()
//
//	// Wait for all goroutines to complete before exiting
//	wg.Wait()
//}
//
//func start() error {
//	env, err := evictedconf.GetEnvVars()
//	if err != nil {
//		return err
//	}
//
//	if env.Type == "webpage" {
//		err := watchPage(env.Endpoint, env.Interval)
//		if err != nil {
//			return err
//		}
//	} else if env.Type == "ip" {
//		watchIp(env.Endpoint, env.Port, env.Interval)
//	} else {
//		return errors.New("Unsupported watch type: " + env.Type)
//	}
//	return nil
//}

func main() {
	log.Println("Work in Progress!!")
	http.HandleFunc("/healthz", healthz.HealthCheck)
	http.ListenAndServe(":8080", nil)
}
