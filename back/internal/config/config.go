package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	// HTTP
	HttpPrefix string
	HttpPort   int

	// DB
	DBHost     string
	DBPort     int
	DBUser     string
	DBPassword string
	DBName     string
}

func InitConfig() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		return nil, fmt.Errorf("error loading .env file: %v", err)
	}

	httpPrefix := os.Getenv("HTTP_PREFIX")
	if httpPrefix == "" {
		return nil, fmt.Errorf("prefix is not set")
	}

	httpPortString := os.Getenv("HTTP_PORT")
	if httpPortString == "" {
		return nil, fmt.Errorf("port is not set")
	}

	httpPort, err := strconv.Atoi(httpPortString)
	if err != nil {
		return nil, fmt.Errorf("error converting port to int: %v", err)
	}

	dbHost := os.Getenv("DB_HOST")
	if dbHost == "" {
		return nil, fmt.Errorf("db host is not set")
	}

	dbPortString := os.Getenv("DB_PORT")
	if dbPortString == "" {
		return nil, fmt.Errorf("db port is not set")
	}

	dbPort, err := strconv.Atoi(dbPortString)
	if err != nil {
		return nil, fmt.Errorf("error converting port to int: %v", err)
	}

	dbUser := os.Getenv("DB_USER")
	if dbUser == "" {
		return nil, fmt.Errorf("db user is not set")
	}

	dbPassword := os.Getenv("DB_PASSWORD")
	if dbPassword == "" {
		return nil, fmt.Errorf("db password is not set")
	}

	dbName := os.Getenv("DB_NAME")
	if dbName == "" {
		return nil, fmt.Errorf("db name is not set")
	}

	return &Config{
		HttpPrefix: httpPrefix,
		HttpPort:   httpPort,
		DBHost:     dbHost,
		DBPort:     dbPort,
		DBUser:     dbUser,
		DBPassword: dbPassword,
		DBName:     dbName,
	}, nil
}
