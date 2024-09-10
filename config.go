package main

import (
	"fmt"
	"os"
)

type Config struct {
	APIKey      string
	APIEndpoint string
	Units       string
}

func loadConfig() (*Config, error) {
	apiKey := os.Getenv("WEATHER_API_KEY")
	if apiKey == "" {
		return nil, fmt.Errorf("WEATHER_API_KEY environment variable is required")
	}

	return &Config{
		APIKey:      apiKey,
		APIEndpoint: "http://api.weatherapi.com/v1/current.json",
		Units:       "metric",
	}, nil
}

func validateCity(city string) error {
	if city == "" {
		return fmt.Errorf("city name cannot be empty")
	}
	
	// Simple validation - you could add more complex validation here
	if len(city) < 2 {
		return fmt.Errorf("city name too short")
	}
	
	return nil
}