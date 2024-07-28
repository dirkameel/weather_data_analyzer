package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

const (
	apiKey      = "your_api_key_here" // Replace with your WeatherAPI.com key
	baseURL     = "http://api.weatherapi.com/v1"
	forecastDays = 3
)

func fetchWeatherData(location string) (*WeatherData, error) {
	// Encode the location for URL
	encodedLocation := url.QueryEscape(location)
	
	// Build the API URL
	apiURL := fmt.Sprintf("%s/forecast.json?key=%s&q=%s&days=%d&aqi=no&alerts=no",
		baseURL, apiKey, encodedLocation, forecastDays)

	// Make HTTP request
	resp, err := http.Get(apiURL)
	if err != nil {
		return nil, fmt.Errorf("failed to make API request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("API returned status %d: %s", resp.StatusCode, string(body))
	}

	// Parse JSON response
	var weatherData WeatherData
	if err := json.NewDecoder(resp.Body).Decode(&weatherData); err != nil {
		return nil, fmt.Errorf("failed to parse JSON response: %v", err)
	}

	return &weatherData, nil
}