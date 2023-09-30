package evictedconf

import (
	"errors"
	"os"
	"regexp"
	"strings"
)

type EnvVars struct {
	Type     string
	Endpoint string
	Interval string
	Port     string
}

func GetEnvVars() (*EnvVars, error) {

	envVars := &EnvVars{
		Type:     os.Getenv("TYPE"),
		Endpoint: os.Getenv("ENDPOINT"),
		Interval: os.Getenv("INTERVAL"),
		Port:     os.Getenv("PORT"),
	}

	if envVars.Type == "" {
		return nil, errors.New("TYPE environment variable is not set")
	}

	if envVars.Type != "webpage" && envVars.Type != "ip" {
		return nil, errors.New("TYPE must be 'webpage' or 'ip'")
	}

	if envVars.Type == "webpage" {
		if envVars.Endpoint == "" {
			return nil, errors.New("ENDPOINT environment variable is not set")
		}
		if !isValidURL(envVars.Endpoint) {
			return nil, errors.New("ENDPOINT must be a valid URL")
		}
	} else if envVars.Type == "ip" {
		if envVars.Endpoint == "" {
			return nil, errors.New("ENDPOINT environment variable is not set")
		}
		if !isValidIP(envVars.Endpoint) {
			return nil, errors.New("ENDPOINT must be a valid IP address")
		}
		if envVars.Port == "" {
			return nil, errors.New("PORT environment variable is not set for 'ip' type")
		}
	}

	if strings.HasSuffix(envVars.Interval, "s") {
		envVars.Interval = strings.TrimSuffix(envVars.Interval, "s")
	}

	if envVars.Interval == "" {
		envVars.Interval = "60"
	}

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
