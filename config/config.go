package config

import (
	"os"
	"strconv"
)

const (
	EnvGoClassApiPort = "GOCLASS_PORT"
	EnvGoClassApiENV  = "GOCLASS_ENV"
	EnvGoClassLogPort = "GOCLASS_LOG_PORT"
	EnvGoClassLogHost = "GOCLASS_LOG_HOST"
)

type Configuration struct {
	Environment string
	Server      ServerConfiguration
	Log         LogConfiguration
}

type ServerConfiguration struct {
	Port int
}

type LogConfiguration struct {
	Host string
	Port int
}

func (c *Configuration) Load() {

	server_config := ServerConfiguration{
		Port: getEnvInt(EnvGoClassApiPort, 8080),
	}

	log_config := LogConfiguration{
		Host: getEnvString(EnvGoClassLogHost, "127.0.0.1"),
		Port: getEnvInt(EnvGoClassLogPort, 24224),
	}

	c.Server = server_config
	c.Log = log_config
	c.Environment = getEnvString(EnvGoClassApiENV, "development")
}

func getEnvString(ev, dv string) string {
	c := os.Getenv(ev)
	if c == "" {
		return dv
	}
	return c
}

func getEnvInt(ev string, dv int) int {
	c := os.Getenv(ev)
	v, err := strconv.Atoi(c)

	if err != nil {
		v = dv
	}

	return v
}
