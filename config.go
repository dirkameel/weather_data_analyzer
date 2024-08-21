package main

import (
	"os"
)

// Config holds application configuration
type Config struct {
	APIKey     string
	APIBaseURL string
	Days       int
}

// NewConfig creates a new configuration instance
func NewConfig() *Config {
	return &Config{
		APIKey:     getEnv("WEATHER_API_KEY", "demo_key"),
		APIBaseURL: "http://api.weatherapi.com/v1",
		Days:       7,
	}
}

// getEnv gets environment variable with fallback
func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}