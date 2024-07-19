package config

import (
	"log"
	"os"
	"strconv"
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

func LoadConfig() Config {
	port, err := strconv.Atoi(os.Getenv("DB_PORT"))
	if err != nil {
		log.Fatalf("Invalid port number: %v", err)
	}

	return Config{
		Server: ServerConfig{
			Address: ":8080",
		},
		Database: DatabaseConfig{
			Host:     os.Getenv("DB_HOST"),
			Port:     port,
			User:     os.Getenv("DB_USER"),
			Password: os.Getenv("DB_PASSWORD"),
			DBName:   os.Getenv("DB_NAME"),
		},
	}
}
