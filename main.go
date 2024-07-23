package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

type WeatherData struct {
	Location struct {
		Name    string `json:"name"`
		Country string `json:"country"`
	} `json:"location"`
	Current struct {
		TempC float64 `json:"temp_c"`
		TempF float64 `json:"temp_f"`
	} `json:"current"`
	Forecast struct {
		Forecastday []ForecastDay `json:"forecastday"`
	} `json:"forecast"`
}

type ForecastDay struct {
	Date string `json:"date"`
	Day  struct {
		MaxTempC float64 `json:"maxtemp_c"`
		MinTempC float64 `json:"mintemp_c"`
		AvgTempC float64 `json:"avgtemp_c"`
	} `json:"day"`
	Hour []HourData `json:"hour"`
}

type HourData struct {
	TimeEpoch int64   `json:"time_epoch"`
	TempC     float64 `json:"temp_c"`
	TempF     float64 `json:"temp_f"`
}

type WeatherAnalyzer struct {
	data     *WeatherData
	location string
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <location>")
		fmt.Println("Example: go run main.go London")
		return
	}

	location := strings.Join(os.Args[1:], " ")
	
	fmt.Printf("🌤️  Fetching weather data for: %s\n", location)
	
	weatherData, err := fetchWeatherData(location)
	if err != nil {
		log.Fatalf("Error fetching weather data: %v", err)
	}

	analyzer := NewWeatherAnalyzer(weatherData, location)
	analyzer.DisplayCurrentWeather()
	analyzer.AnalyzeTrends()
	analyzer.GenerateVisualization()
}