package watcher

import (
	"errors"
	"github.com/timam/statuz/config"
)

func Start() error {
	env, err := config.GetEnvVars()
	if err != nil {
		return err
	}

	if env.Type == "webpage" {
		return watchPage(env.Endpoint, env.Interval)
	} else if env.Type == "ip" {
		return watchIp(env.Endpoint, env.Port, env.Interval)
	} else {
		return errors.New("Unsupported watch type: " + env.Type)
	}
}
