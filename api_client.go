package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

const (
	apiKey      = "your_api_key_here" // Replace with your OpenWeatherMap API key
	baseURL     = "https://api.openweathermap.org/data/2.5/weather"
	tempUnit    = "metric" // Use "imperial" for Fahrenheit
)

func fetchCurrentWeather(city string) (WeatherData, error) {
	url := fmt.Sprintf("%s?q=%s&appid=%s&units=%s", baseURL, city, apiKey, tempUnit)
	
	resp, err := http.Get(url)
	if err != nil {
		return WeatherData{}, fmt.Errorf("HTTP request failed: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return WeatherData{}, fmt.Errorf("API error: %s - %s", resp.Status, string(body))
	}

	var weatherResp WeatherResponse
	if err := json.NewDecoder(resp.Body).Decode(&weatherResp); err != nil {
		return WeatherData{}, fmt.Errorf("JSON decode failed: %v", err)
	}

	return WeatherData{
		City:      weatherResp.Name,
		Temp:      weatherResp.Main.Temp,
		Humidity:  weatherResp.Main.Humidity,
		Timestamp: time.Now(),
	}, nil
}

func saveWeatherData(data []WeatherData) error {
	file, err := os.Create("weather_data.json")
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	return encoder.Encode(data)
}