package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

type WeatherData struct {
	Location struct {
		Name    string `json:"name"`
		Country string `json:"country"`
	} `json:"location"`
	Current struct {
		TempC     float64 `json:"temp_c"`
		Condition struct {
			Text string `json:"text"`
		} `json:"condition"`
		Humidity   int     `json:"humidity"`
		WindKph    float64 `json:"wind_kph"`
		FeelsLikeC float64 `json:"feelslike_c"`
	} `json:"current"`
	Forecast struct {
		Forecastday []struct {
			Date string `json:"date"`
			Day  struct {
				MaxTempC float64 `json:"maxtemp_c"`
				MinTempC float64 `json:"mintemp_c"`
				AvgTempC float64 `json:"avgtemp_c"`
			} `json:"day"`
		} `json:"forecastday"`
	} `json:"forecast"`
}

type WeatherAnalyzer struct {
	Data []WeatherData
}

func main() {
	fmt.Println("🌤️  Weather Data Analyzer")
	fmt.Println("==========================")

	// Get location from user or use default
	location := getLocationInput()

	// Fetch weather data
	weatherData, err := fetchWeatherData(location)
	if err != nil {
		log.Fatalf("Error fetching weather data: %v", err)
	}

	// Analyze and display results
	analyzer := &WeatherAnalyzer{Data: []WeatherData{weatherData}}
	analyzer.DisplayCurrentWeather()
	analyzer.AnalyzeTemperatureTrends()
	analyzer.VisualizeTemperatureTrends()
}

func getLocationInput() string {
	if len(os.Args) > 1 {
		return os.Args[1]
	}
	
	fmt.Print("Enter location (or press Enter for London): ")
	var location string
	fmt.Scanln(&location)
	
	if location == "" {
		return "London"
	}
	return location
}

func fetchWeatherData(location string) (WeatherData, error) {
	apiKey := getAPIKey()
	url := fmt.Sprintf("http://api.weatherapi.com/v1/forecast.json?key=%s&q=%s&days=7&aqi=no&alerts=no", apiKey, location)

	resp, err := http.Get(url)
	if err != nil {
		return WeatherData{}, fmt.Errorf("failed to fetch data: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return WeatherData{}, fmt.Errorf("API returned status: %s", resp.Status)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return WeatherData{}, fmt.Errorf("failed to read response: %v", err)
	}

	var weatherData WeatherData
	err = json.Unmarshal(body, &weatherData)
	if err != nil {
		return WeatherData{}, fmt.Errorf("failed to parse JSON: %v", err)
	}

	return weatherData, nil
}

func getAPIKey() string {
	// First try environment variable
	if apiKey := os.Getenv("WEATHER_API_KEY"); apiKey != "" {
		return apiKey
	}

	// Then try config file
	config, err := loadConfig()
	if err == nil && config.APIKey != "" {
		return config.APIKey
	}

	// Fallback to demo key (limited usage)
	fmt.Println("⚠️  Using demo API key (limited requests)")
	return "demo_key_will_not_work_use_real_key"
}