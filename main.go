package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
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
		TempC     float64 `json:"temp_c"`
		Condition struct {
			Text string `json:"text"`
		} `json:"condition"`
		Humidity   float64 `json:"humidity"`
		WindKph    float64 `json:"wind_kph"`
		FeelsLikeC float64 `json:"feelslike_c"`
	} `json:"current"`
	Forecast struct {
		Forecastday []struct {
			Date string `json:"date"`
			Day  struct {
				MaxtempC float64 `json:"maxtemp_c"`
				MintempC float64 `json:"mintemp_c"`
				AvgtempC float64 `json:"avgtemp_c"`
			} `json:"day"`
		} `json:"forecastday"`
	} `json:"forecast"`
}

type WeatherAnalyzer struct {
	Data []WeatherData
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <city>")
		fmt.Println("Example: go run main.go London")
		return
	}

	city := os.Args[1]
	
	// Fetch weather data
	weatherData, err := fetchWeatherData(city)
	if err != nil {
		log.Fatal("Error fetching weather data:", err)
	}

	// Analyze and display results
	analyzer := &WeatherAnalyzer{Data: []WeatherData{*weatherData}}
	analyzer.DisplayCurrentWeather()
	analyzer.DisplayTemperatureTrends()
	analyzer.GenerateVisualization()
}

func fetchWeatherData(city string) (*WeatherData, error) {
	apiKey := "your_api_key_here" // Replace with actual API key
	url := fmt.Sprintf("http://api.weatherapi.com/v1/forecast.json?key=%s&q=%s&days=3&aqi=no&alerts=no", apiKey, city)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var weatherData WeatherData
	err = json.Unmarshal(body, &weatherData)
	if err != nil {
		return nil, err
	}

	return &weatherData, nil
}

func (wa *WeatherAnalyzer) DisplayCurrentWeather() {
	if len(wa.Data) == 0 {
		fmt.Println("No weather data available")
		return
	}

	data := wa.Data[0]
	fmt.Println("\n=== CURRENT WEATHER ===")
	fmt.Printf("Location: %s, %s\n", data.Location.Name, data.Location.Country)
	fmt.Printf("Temperature: %.1f°C (Feels like: %.1f°C)\n", data.Current.TempC, data.Current.FeelsLikeC)
	fmt.Printf("Condition: %s\n", data.Current.Condition.Text)
	fmt.Printf("Humidity: %.0f%%\n", data.Current.Humidity)
	fmt.Printf("Wind: %.1f km/h\n", data.Current.WindKph)
}

func (wa *WeatherAnalyzer) DisplayTemperatureTrends() {
	if len(wa.Data) == 0 {
		return
	}

	data := wa.Data[0]
	fmt.Println("\n=== TEMPERATURE TRENDS (3-Day Forecast) ===")
	
	for i, day := range data.Forecast.Forecastday {
		date, _ := time.Parse("2006-01-02", day.Date)
		dayName := date.Format("Monday")
		
		fmt.Printf("%s (%s):\n", dayName, day.Date)
		fmt.Printf("  Max: %.1f°C | Min: %.1f°C | Avg: %.1f°C\n", 
			day.Day.MaxtempC, day.Day.MintempC, day.Day.AvgtempC)
		
		if i < len(data.Forecast.Forecastday)-1 {
			nextDay := data.Forecast.Forecastday[i+1]
			tempChange := nextDay.Day.AvgtempC - day.Day.AvgtempC
			trend := "stable"
			if tempChange > 1 {
				trend = "warming"
			} else if tempChange < -1 {
				trend = "cooling"
			}
			fmt.Printf("  Trend: %s (Δ%.1f°C)\n", trend, tempChange)
		}
	}
}

func (wa *WeatherAnalyzer) GenerateVisualization() {
	if len(wa.Data) == 0 {
		return
	}

	data := wa.Data[0]
	fmt.Println("\n=== TEMPERATURE VISUALIZATION ===")
	
	for _, day := range data.Forecast.Forecastday {
		date, _ := time.Parse("2006-01-02", day.Date)
		dayName := date.Format("Mon")
		
		// Create a simple bar chart for temperature range
		rangeWidth := int(day.Day.MaxtempC - day.Day.MintempC)
		minBar := strings.Repeat(" ", int(day.Day.MintempC)+10) // Offset for negative temps
		rangeBar := strings.Repeat "▀", rangeWidth)
		
		fmt.Printf("%s: %5.1f°C ", dayName, day.Day.AvgtempC)
		fmt.Printf("[%s%s] (%.1f°C - %.1f°C)\n", minBar, rangeBar, day.Day.MintempC, day.Day.MaxtempC)
	}
	
	// Additional analysis
	fmt.Println("\n=== WEATHER ANALYSIS ===")
	currentTemp := data.Current.TempC
	avgTemp := calculateAverageTemp(data)
	
	fmt.Printf("Current vs Forecast Average: %.1f°C vs %.1f°C\n", currentTemp, avgTemp)
	
	if currentTemp > avgTemp + 2 {
		fmt.Println("📈 Currently warmer than forecast average")
	} else if currentTemp < avgTemp - 2 {
		fmt.Println("📉 Currently cooler than forecast average")
	} else {
		fmt.Println("➡️  Temperature is close to forecast average")
	}
}

func calculateAverageTemp(data WeatherData) float64 {
	total := 0.0
	count := 0
	
	for _, day := range data.Forecast.Forecastday {
		total += day.Day.AvgtempC
		count++
	}
	
	if count > 0 {
		return total / float64(count)
	}
	return 0
}