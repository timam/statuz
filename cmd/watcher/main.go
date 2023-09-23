package watcher

import (
	"github.com/timam/statuz/config"
	"log"
)

func Start() {
	env, err := config.GetEnvVars()
	if err != nil {
		log.Fatalln(err)
	}

	if env.Type == "website" {
		watchWebsite(env.Endpoint, env.Interval)
	}

}
