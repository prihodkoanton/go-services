package config

import (
	"fmt"
	"os"
	"strconv"
)

type Config struct {
	Server struct {
		Port int
	}
	Database struct {
		Host     string
		Port     int
		User     string
		Password string
		Name     string
		SSLMode  string
	}
}

func LoadConfig() (*Config, error) {
	cfg := &Config{}

	// Server
	portStr := getEnv("PORT", "9998")
	port, err := strconv.Atoi(portStr)
	if err != nil {
		return nil, fmt.Errorf("invalid PORT value: %v", err)
	}
	cfg.Server.Port = port

	// Database
	cfg.Database.Host = getEnv("DB_HOST", "localhost")

	dbPortStr := getEnv("DB_PORT", "9999")
	dbPort, err := strconv.Atoi(dbPortStr)
	if err != nil {
		return nil, fmt.Errorf("invalid DB_PORT value: %v", err)
	}
	cfg.Database.Port = dbPort

	cfg.Database.User = getEnv("DB_USER", "postgres")
	cfg.Database.Password = getEnv("DB_PASSWORD", "1234")
	cfg.Database.Name = getEnv("DB_NAME", "postgres")
	cfg.Database.SSLMode = getEnv("DB_SSLMODE", "disable")

	return cfg, nil
}

func getEnv(key, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultVal
}
