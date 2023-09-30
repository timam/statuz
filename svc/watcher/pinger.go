package main

import (
	"fmt"
	"log"
	"net"
	"strconv"
	"time"
)

func watchIp(ip string, port string, intervalInSeconds string) error {
	log.Printf("starting watcher for %s:%s\n", ip, port)

	interval, err := strconv.ParseInt(intervalInSeconds, 10, 64)
	if err != nil {
		return err
	}

	ticker := time.NewTicker(time.Duration(interval) * time.Second)
	defer ticker.Stop()

	for range ticker.C {
		pinger(ip, port)
	}

	return nil
}

func pinger(ip string, port string) {
	address := fmt.Sprintf("%s:%s", ip, port)

	conn, err := net.Dial("tcp", address)
	if err != nil {
		log.Printf("ping failed %s: %v\n", address, err)
		return
	}
	defer conn.Close()

	log.Printf("Successfully pinged %s:%s\n", ip, port)
}
