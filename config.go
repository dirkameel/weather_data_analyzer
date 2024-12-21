package main

import (
	"encoding/json"
	"os"
)

type Config struct {
	APIKey string `json:"api_key"`
	Units  string `json:"units"`
	City   string `json:"default_city"`
}

func loadConfig() (Config, error) {
	var config Config
	
	// Check if config file exists
	if _, err := os.Stat("config.json"); os.IsNotExist(err) {
		return config, err
	}

	file, err := os.ReadFile("config.json")
	if err != nil {
		return config, err
	}

	err = json.Unmarshal(file, &config)
	return config, err
}

func createDefaultConfig() error {
	config := Config{
		APIKey: "YOUR_API_KEY_HERE",
		Units:  "metric",
		City:   "London",
	}

	file, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile("config.json", file, 0644)
}