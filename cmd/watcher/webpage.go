package watcher

import (
	"log"
	"net/http"
	"strconv"
	"time"
)

func watchPage(url string, intervalInSeconds string) {
	log.Printf("starting watcher for %s\n", url)

	interval, _ := strconv.ParseInt(intervalInSeconds, 10, 64)

	ticker := time.NewTicker(time.Duration(interval) * time.Second)

	// Start a goroutine to check the URL status periodically
	go func() {
		for {
			select {
			case <-ticker.C:
				pageChecker(url)
			}
		}
	}()

	// Block here to keep the watcher running
	select {}
}

func pageChecker(url string) {
	resp, err := http.Get(url)
	if err != nil {
		log.Printf("Error checking URL %s: %v\n", url, err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		log.Printf("URL %s is returning 200 OK\n", url)
	} else {
		log.Printf("URL %s is returning status code %d\n", url, resp.StatusCode)
	}
}
