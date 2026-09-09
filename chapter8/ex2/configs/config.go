package configs

import "os"

type Config struct {
	AppName string
	Port    string
	DBHost  string
}

func LoadConfig() Config {

	appName := os.Getenv("APP_NAME")
	if appName == "" {
		appName = "demo-service"
	}

	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8080"
	}

	dbHost := os.Getenv("DB_HOST")
	if dbHost == "" {
		dbHost = "localhost"
	}

	return Config{
		AppName: appName,
		Port:    port,
		DBHost:  dbHost,
	}
}
