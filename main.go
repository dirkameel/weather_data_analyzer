package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

type WeatherConfig struct {
	APIKey string `json:"api_key"`
	City   string `json:"city"`
}

type WeatherResponse struct {
	Main struct {
		Temp     float64 `json:"temp"`
		Humidity int     `json:"humidity"`
	} `json:"main"`
	Name string `json:"name"`
}

type WeatherData struct {
	City      string    `json:"city"`
	Temp      float64   `json:"temp"`
	Humidity  int       `json:"humidity"`
	Timestamp time.Time `json:"timestamp"`
}

func main() {
	// Load configuration
	config, err := loadConfig("config.json")
	if err != nil {
		log.Fatal("Error loading config:", err)
	}

	// Fetch current weather
	weather, err := fetchWeatherData(config)
	if err != nil {
		log.Fatal("Error fetching weather data:", err)
	}

	// Store the data
	err = storeWeatherData(weather)
	if err != nil {
		log.Fatal("Error storing weather data:", err)
	}

	// Analyze and visualize trends
	err = analyzeAndVisualize()
	if err != nil {
		log.Fatal("Error analyzing data:", err)
	}

	fmt.Printf("Weather data processed successfully!\n")
	fmt.Printf("Current temperature in %s: %.1f°C\n", weather.City, weather.Temp)
}