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

	if env.Type == "webpage" {
		watchPage(env.Endpoint, env.Interval)
	}
	if env.Type == "ip" {
		watchIp(env.Endpoint, env.Port, env.Interval)
	}

}
