package config

import (
	"errors"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

type EnvVars struct {
	Type     string
	Endpoint string
	Interval string
}

func GetEnvVars() (*EnvVars, error) {

	envVars := &EnvVars{
		Type:     os.Getenv("TYPE"),
		Endpoint: os.Getenv("ENDPOINT"),
		Interval: os.Getenv("INTERVAL"),
	}

	if envVars.Type == "" {
		return nil, errors.New("TYPE environment variable is not set")
	}

	if envVars.Endpoint == "" {
		return nil, errors.New("ENDPOINT environment variable is not set")
	}

	if envVars.Type != "webpage" && envVars.Type != "ip" {
		return nil, errors.New("TYPE must be 'webpage' or 'ip'")
	}

	if envVars.Type == "webpage" {
		if !isValidURL(envVars.Endpoint) {
			return nil, errors.New("ENDPOINT must be a valid URL")
		}
	}

	if envVars.Type == "ip" {
		if !isValidIP(envVars.Endpoint) {
			return nil, errors.New("ENDPOINT must be a valid IP address")
		}
	}

	if strings.HasSuffix(envVars.Interval, "s") {
		envVars.Interval = strings.TrimSuffix(envVars.Interval, "s")
	}

	if envVars.Interval == "" {
		envVars.Interval = "300"
	}

	printEnvs(envVars)

	return envVars, nil
}

func isValidURL(s string) bool {
	regex := `^(https?):\/\/[^\s/$.?#].[^\s]*$`
	return regexp.MustCompile(regex).MatchString(s)
}

func isValidIP(s string) bool {
	regex := `^(\d{1,3}\.\d{1,3}\.\d{1,3}\.\d{1,3})$`
	return regexp.MustCompile(regex).MatchString(s)
}

func printEnvs(env *EnvVars) {
	message := fmt.Sprintf("Type: %s, Endpoint: %s, Interval: %s", env.Type, env.Endpoint, env.Interval)
	log.Println(message)
}
