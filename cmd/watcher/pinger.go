package watcher

import (
	"fmt"
	"log"
	"net"
	"strconv"
	"time"
)

func watchIp(ip string, port string, intervalInSeconds string) {
	log.Printf("starting watcher for %s:%s\n", ip, port)

	interval, _ := strconv.ParseUint(intervalInSeconds, 10, 64)

	ticker := time.NewTicker(time.Duration(interval) * time.Second)

	go func() {
		for {
			select {
			case <-ticker.C:
				pinger(ip, port)
			}
		}
	}()
}

func pinger(ip string, port string) {
	address := fmt.Sprintf("%s:%s", ip, port)

	conn, err := net.Dial("tcp", address)
	if err != nil {
		log.Printf("ping failed %s\n", address) // Connection failed
	}
	defer conn.Close()
	fmt.Printf("Successfully pinged %s:%s\n", ip, port)

}
