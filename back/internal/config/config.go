package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	HttpServer string
	HttpPrefix string
	DBPATH     string
}

func InitConfig() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		return nil, fmt.Errorf("error loading .env file: %v", err)
	}

	port := os.Getenv("HTTP_SERVER_PORT")
	if port == "" {
		port = "10000"
	}

	baseApiPrefix := os.Getenv("BASE_API_PREFIX")
	if baseApiPrefix == "" {
		return nil, fmt.Errorf("prefix is not set")
	}

	databaseURL := os.Getenv("DATABASE_PATH")
	if databaseURL == "" {
		return nil, fmt.Errorf("databaseURL is not set")
	}

	return &Config{
		HttpServer: port,
		HttpPrefix: baseApiPrefix,
		DBPATH:     databaseURL,
	}, nil
}
