package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

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
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <city1> <city2> ...")
		fmt.Println("Example: go run main.go London Tokyo NewYork")
		os.Exit(1)
	}

	cities := os.Args[1:]
	fmt.Printf("Fetching weather data for: %v\n", cities)

	// Fetch current weather data
	var weatherData []WeatherData
	for _, city := range cities {
		data, err := fetchCurrentWeather(city)
		if err != nil {
			log.Printf("Error fetching data for %s: %v", city, err)
			continue
		}
		weatherData = append(weatherData, data)
	}

	// Save data to file
	if err := saveWeatherData(weatherData); err != nil {
		log.Fatalf("Error saving weather data: %v", err)
	}

	// Analyze and visualize
	analyzeWeatherData(weatherData)
	generateVisualization(weatherData)
}