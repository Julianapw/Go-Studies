package config

import "os"

type Config struct {
	Port string
	Name string
}

func Load() Config {

	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8080"
	}

	name := os.Getenv("APP_NAME")
	if name == "" {
		name = "task-service"
	}

	return Config{
		Port: port,
		Name: name,
	}
}
