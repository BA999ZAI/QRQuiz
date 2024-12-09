package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	HttpServer string
	HttpPrefix string
	DBURL      string
}

func InitConfig() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		return nil, fmt.Errorf("error loading .env file: %v", err)
	}

	port := os.Getenv("HTTP_SERVER_PORT")
	if port == "" {
		return nil, fmt.Errorf("Port is not set")
	}

	baseApiPrefix := os.Getenv("BASE_API_PREFIX")
	if baseApiPrefix == "" {
		return nil, fmt.Errorf("Prefix is not set")
	}

	databaseURL := os.Getenv("DB_URL")
	if databaseURL == "" {
		return nil, fmt.Errorf("DatabaseURL is not set")
	}

	return &Config{
		HttpServer: port,
		HttpPrefix: baseApiPrefix,
	}, nil
}
