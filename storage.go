package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

const dataFile = "weather_data.json"

func storeWeatherData(weather *WeatherData) error {
	// Load existing data
	var allData []WeatherData
	file, err := os.Open(dataFile)
	if err == nil {
		json.NewDecoder(file).Decode(&allData)
		file.Close()
	}

	// Add new data
	allData = append(allData, *weather)

	// Keep only last 24 hours of data
	allData = filterRecentData(allData)

	// Save back to file
	file, err = os.Create(dataFile)
	if err != nil {
		return fmt.Errorf("could not create data file: %v", err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	return encoder.Encode(allData)
}

func filterRecentData(data []WeatherData) []WeatherData {
	var recent []WeatherData
	cutoff := time.Now().Add(-24 * time.Hour)
	
	for _, item := range data {
		if item.Timestamp.After(cutoff) {
			recent = append(recent, item)
		}
	}
	return recent
}

func loadWeatherData() ([]WeatherData, error) {
	var data []WeatherData
	file, err := os.Open(dataFile)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	err = json.NewDecoder(file).Decode(&data)
	return data, err
}