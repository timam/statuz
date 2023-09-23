package watcher

import (
	"github.com/timam/statuz/config"
	"log"
)

func Start() {
	_, err := config.GetEnvVars()
	if err != nil {
		log.Fatalln(err)
	}
}
