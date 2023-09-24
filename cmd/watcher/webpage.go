package watcher

import (
	"log"
	"net/http"
	"strconv"
	"time"
)

func watchPage(url string, intervalInSeconds string) error {
	log.Printf("starting watcher for %s\n", url)

	interval, err := strconv.ParseInt(intervalInSeconds, 10, 64)
	if err != nil {
		return err
	}

	ticker := time.NewTicker(time.Duration(interval) * time.Second)
	defer ticker.Stop()

	for range ticker.C {
		pageChecker(url)
	}

	return nil
}


func pageChecker(url string) {
	resp, err := http.Get(url)
	if err != nil {
		log.Printf("Error checking URL %s: %v\n", url, err)
		return
	}
	defer resp.Body.Close()

	statusText := http.StatusText(resp.StatusCode)
	log.Printf("URL %s is returning %s\n", url, statusText)
}

