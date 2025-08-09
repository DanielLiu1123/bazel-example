package config

import (
	"os"
)

// Config holds the application configuration
type Config struct {
	Port            string
	UserServiceURL  string
	OrderServiceURL string
}

// Load loads configuration from environment variables
func Load() *Config {
	return &Config{
		Port:            getEnv("PORT", "8080"),
		UserServiceURL:  getEnv("USER_SERVICE_URL", "http://localhost:8081"),
		OrderServiceURL: getEnv("ORDER_SERVICE_URL", "http://localhost:8082"),
	}
}

// getEnv gets an environment variable with a default value
func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
