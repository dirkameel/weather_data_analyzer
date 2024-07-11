package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

func loadConfig(filename string) (*WeatherConfig, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("could not open config file: %v", err)
	}
	defer file.Close()

	var config WeatherConfig
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&config)
	if err != nil {
		return nil, fmt.Errorf("could not decode config: %v", err)
	}

	return &config, nil
}

func fetchWeatherData(config *WeatherConfig) (*WeatherData, error) {
	url := fmt.Sprintf("http://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s&units=metric",
		config.City, config.APIKey)

	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("HTTP request failed: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("API error: %s - %s", resp.Status, string(body))
	}

	var weatherResp WeatherResponse
	err = json.NewDecoder(resp.Body).Decode(&weatherResp)
	if err != nil {
		return nil, fmt.Errorf("could not decode weather response: %v", err)
	}

	return &WeatherData{
		City:      weatherResp.Name,
		Temp:      weatherResp.Main.Temp,
		Humidity:  weatherResp.Main.Humidity,
		Timestamp: time.Now(),
	}, nil
}