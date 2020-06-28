package config

import (
	"os"
	"strconv"
)

type Configuration struct {
	Environment string
	Port        int
}

func (c *Configuration) Load() {
	port, p_err := strconv.Atoi(getConfigOrDefault("API_PORT", "8080"))

	if p_err != nil {
		port = 8080
	}

	c.Port = port
	c.Environment = getConfigOrDefault("ENVIRONMENT", "development")
}

func getConfigOrDefault(ev, dv string) string {
	c := os.Getenv(ev)
	if c == "" {
		return dv
	}
	return c
}
