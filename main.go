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

type WeatherConfig struct {
	APIKey string `json:"api_key"`
	City   string `json:"city"`
	Units  string `json:"units"`
}

type WeatherData struct {
	Main struct {
		Temp     float64 `json:"temp"`
		Humidity int     `json:"humidity"`
		Pressure int     `json:"pressure"`
	} `json:"main"`
	Weather []struct {
		Description string `json:"description"`
		Main        string `json:"main"`
	} `json:"weather"`
	Wind struct {
		Speed float64 `json:"speed"`
	} `json:"wind"`
	Name string `json:"name"`
	DT   int64  `json:"dt"`
}

type ForecastData struct {
	List []struct {
		DT   int64 `json:"dt"`
		Main struct {
			Temp float64 `json:"temp"`
		} `json:"main"`
		Weather []struct {
			Description string `json:"description"`
		} `json:"weather"`
	} `json:"list"`
	City struct {
		Name string `json:"name"`
	} `json:"city"`
}

func main() {
	// Load configuration
	config, err := loadConfig("config.json")
	if err != nil {
		log.Fatal("Error loading config:", err)
	}

	fmt.Println("🌤️  Weather Data Analyzer")
	fmt.Println("==========================")

	// Fetch current weather
	currentWeather, err := fetchCurrentWeather(config)
	if err != nil {
		log.Fatal("Error fetching current weather:", err)
	}

	// Fetch forecast
	forecast, err := fetchForecast(config)
	if err != nil {
		log.Fatal("Error fetching forecast:", err)
	}

	// Display current weather
	displayCurrentWeather(currentWeather)

	// Analyze and visualize trends
	analyzeTrends(forecast)

	// Generate visualization
	generateVisualization(forecast)
}

func loadConfig(filename string) (*WeatherConfig, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var config WeatherConfig
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}

func fetchCurrentWeather(config *WeatherConfig) (*WeatherData, error) {
	url := fmt.Sprintf(
		"https://api.openweathermap.org/data/2.5/weather?q=%s&units=%s&appid=%s",
		config.City, config.Units, config.APIKey,
	)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

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

func fetchForecast(config *WeatherConfig) (*ForecastData, error) {
	url := fmt.Sprintf(
		"https://api.openweathermap.org/data/2.5/forecast?q=%s&units=%s&appid=%s",
		config.City, config.Units, config.APIKey,
	)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var forecast ForecastData
	err = json.Unmarshal(body, &forecast)
	if err != nil {
		return nil, err
	}

	return &forecast, nil
}

func displayCurrentWeather(weather *WeatherData) {
	fmt.Printf("\n📍 Current Weather in %s\n", weather.Name)
	fmt.Printf("🌡️  Temperature: %.1f°C\n", weather.Main.Temp)
	fmt.Printf("💧 Humidity: %d%%\n", weather.Main.Humidity)
	fmt.Printf("🌬️  Wind Speed: %.1f m/s\n", weather.Wind.Speed)
	fmt.Printf("📊 Pressure: %d hPa\n", weather.Main.Pressure)
	
	if len(weather.Weather) > 0 {
		fmt.Printf("☁️  Conditions: %s (%s)\n", 
			weather.Weather[0].Description, weather.Weather[0].Main)
	}
	
	timestamp := time.Unix(weather.DT, 0)
	fmt.Printf("🕒 Last Updated: %s\n", timestamp.Format("2006-01-02 15:04:05"))
}

func analyzeTrends(forecast *ForecastData) {
	fmt.Printf("\n📈 Temperature Trend Analysis for %s\n", forecast.City.Name)
	fmt.Println("=====================================")

	if len(forecast.List) == 0 {
		fmt.Println("No forecast data available")
		return
	}

	// Calculate statistics
	minTemp := forecast.List[0].Main.Temp
	maxTemp := forecast.List[0].Main.Temp
	sumTemp := 0.0

	for i, item := range forecast.List {
		if item.Main.Temp < minTemp {
			minTemp = item.Main.Temp
		}
		if item.Main.Temp > maxTemp {
			maxTemp = item.Main.Temp
		}
		sumTemp += item.Main.Temp

		// Display next 5 data points
		if i < 5 {
			timestamp := time.Unix(item.DT, 0)
			fmt.Printf("📅 %s: %.1f°C - %s\n", 
				timestamp.Format("01/02 15:04"), 
				item.Main.Temp, 
				item.Weather[0].Description)
		}
	}

	avgTemp := sumTemp / float64(len(forecast.List))
	fmt.Printf("\n📊 Statistics:\n")
	fmt.Printf("   Minimum Temperature: %.1f°C\n", minTemp)
	fmt.Printf("   Maximum Temperature: %.1f°C\n", maxTemp)
	fmt.Printf("   Average Temperature: %.1f°C\n", avgTemp)
	fmt.Printf("   Temperature Range: %.1f°C\n", maxTemp-minTemp)
}

func generateVisualization(forecast *ForecastData) {
	fmt.Printf("\n📊 Temperature Visualization\n")
	fmt.Println("============================")

	if len(forecast.List) == 0 {
		return
	}

	// Find min and max for scaling
	minTemp := forecast.List[0].Main.Temp
	maxTemp := forecast.List[0].Main.Temp
	for _, item := range forecast.List {
		if item.Main.Temp < minTemp {
			minTemp = item.Main.Temp
		}
		if item.Main.Temp > maxTemp {
			maxTemp = item.Main.Temp
		}
	}

	// Simple ASCII visualization
	tempRange := maxTemp - minTemp
	if tempRange == 0 {
		tempRange = 1 // Avoid division by zero
	}

	fmt.Println("Legend: █ = 2°C increment")
	fmt.Println("Temperature Scale:")

	for i, item := range forecast.List {
		if i >= 8 { // Limit to first 8 data points for readability
			break
		}
		
		timestamp := time.Unix(item.DT, 0)
		normalizedTemp := (item.Main.Temp - minTemp) / tempRange
		bars := int(normalizedTemp * 20) // Scale to 20 characters max
		
		visualization := ""
		for j := 0; j < bars; j++ {
			visualization += "█"
		}
		
		fmt.Printf("%s: %6.1f°C %s\n", 
			timestamp.Format("01/02 15:04"), 
			item.Main.Temp, 
			visualization)
	}

	// Print scale reference
	fmt.Printf("\nScale Reference: %.1f°C to %.1f°C\n", minTemp, maxTemp)
}