package config

import (
	"log"
	"os"
)

type Config struct {
	Server   ServerConfig
	Database DatabaseConfig
}

type ServerConfig struct {
	Address string
}

type DatabaseConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
}

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		if defaultValue == "" {
			log.Fatalf("Environment variable %s not set", key)
		}
		return defaultValue
	}
	return value
}

func LoadConfig() Config {
	// portStr := getEnv("DB_PORT", "")
	// port, err := strconv.Atoi(portStr)
	// if err != nil {
	// 	log.Fatalf("Invalid port number: %v", err)
	// }

	return Config{
		Server: ServerConfig{
			Address: getEnv("SERVER_ADDRESS", ":8080"),
		},
		Database: DatabaseConfig{
			Host:     getEnv("DB_HOST", ""),
			Port:     port,
			User:     getEnv("DB_USER", ""),
			Password: getEnv("DB_PASSWORD", ""),
			DBName:   getEnv("DB_NAME", ""),
		},
	}
}
