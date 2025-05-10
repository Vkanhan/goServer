package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	ServerPort  string
	DBHost      string
	DBPort      string
	DBUser      string
	DBPassword  string
	DBName      string
	DBSSLMode   string
}

func Load() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		fmt.Println("Warning: .env file not found. Using environment variables.")
	}

	return &Config{
		ServerPort:  getEnv("SERVER_PORT"),
		DBHost:      getEnv("DB_HOST"),
		DBPort:      getEnv("DB_PORT"),
		DBUser:      getEnv("DB_USER"),
		DBPassword:  getEnv("DB_PASSWORD"),
		DBName:      getEnv("DB_NAME"),
		DBSSLMode:   getEnv("DB_SSL_MODE"),
	}, nil
}

func getEnv(key string) (string) {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return "" 
}