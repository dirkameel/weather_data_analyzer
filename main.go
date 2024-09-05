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

type WeatherData struct {
	Location struct {
		Name    string `json:"name"`
		Country string `json:"country"`
	} `json:"location"`
	Current struct {
		TempC float64 `json:"temp_c"`
		TempF float64 `json:"temp_f"`
	} `json:"current"`
}

type WeatherAnalysis struct {
	Location    string
	Temperature float64
	Timestamp   time.Time
	Trend       string
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <city_name>")
		fmt.Println("Example: go run main.go London")
		return
	}

	city := os.Args[1]
	
	// Fetch current weather
	weather, err := fetchWeather(city)
	if err != nil {
		log.Fatalf("Error fetching weather: %v", err)
	}

	// Analyze and display results
	analysis := analyzeWeather(weather)
	displayResults(analysis)
	
	// Generate visualization
	generateVisualization(analysis)
}

func fetchWeather(city string) (*WeatherData, error) {
	apiKey := os.Getenv("WEATHER_API_KEY")
	if apiKey == "" {
		return nil, fmt.Errorf("WEATHER_API_KEY environment variable not set")
	}

	url := fmt.Sprintf("http://api.weatherapi.com/v1/current.json?key=%s&q=%s&aqi=no", apiKey, city)
	
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API request failed with status: %s", resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var weather WeatherData
	err = json.Unmarshal(body, &weather)
	if err != nil {
		return nil, err
	}

	return &weather, nil
}

func analyzeWeather(weather *WeatherData) *WeatherAnalysis {
	analysis := &WeatherAnalysis{
		Location:    fmt.Sprintf("%s, %s", weather.Location.Name, weather.Location.Country),
		Temperature: weather.Current.TempC,
		Timestamp:   time.Now(),
	}

	// Simple trend analysis based on temperature
	if weather.Current.TempC < 0 {
		analysis.Trend = "❄️ Freezing"
	} else if weather.Current.TempC < 10 {
		analysis.Trend = "🥶 Cold"
	} else if weather.Current.TempC < 20 {
		analysis.Trend = "😊 Mild"
	} else if weather.Current.TempC < 30 {
		analysis.Trend = "☀️ Warm"
	} else {
		analysis.Trend = "🔥 Hot"
	}

	return analysis
}

func displayResults(analysis *WeatherAnalysis) {
	fmt.Println("\n🌤️  WEATHER ANALYSIS REPORT")
	fmt.Println("============================")
	fmt.Printf("Location:    %s\n", analysis.Location)
	fmt.Printf("Temperature: %.1f°C (%.1f°F)\n", analysis.Temperature, analysis.Temperature*9/5+32)
	fmt.Printf("Condition:   %s\n", analysis.Trend)
	fmt.Printf("Time:        %s\n", analysis.Timestamp.Format("2006-01-02 15:04:05"))
	fmt.Println("============================")
}