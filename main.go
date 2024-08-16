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
		Humidity   float64 `json:"humidity"`
		WindKph    float64 `json:"wind_kph"`
		FeelsLikeC float64 `json:"feelslike_c"`
	} `json:"current"`
	Forecast struct {
		ForecastDay []struct {
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

	if len(os.Args) < 2 {
		log.Fatal("Usage: go run main.go <city_name>")
	}

	city := os.Args[1]
	apiKey := getAPIKey()

	// Fetch current weather and forecast
	weatherData, err := fetchWeatherData(city, apiKey)
	if err != nil {
		log.Fatalf("Error fetching weather data: %v", err)
	}

	// Analyze and display results
	analyzer := &WeatherAnalyzer{Data: []WeatherData{*weatherData}}
	analyzer.DisplayCurrentWeather()
	analyzer.AnalyzeTemperatureTrends()
	analyzer.GenerateVisualization()
}

func getAPIKey() string {
	// In production, use environment variables
	// For demo purposes, you can set this directly or use env var
	apiKey := os.Getenv("WEATHER_API_KEY")
	if apiKey == "" {
		// Using a demo key - replace with your actual API key
		apiKey = "your_weather_api_key_here"
		fmt.Println("⚠️  Using demo mode. Set WEATHER_API_KEY environment variable for real data.")
	}
	return apiKey
}

func fetchWeatherData(city, apiKey string) (*WeatherData, error) {
	url := fmt.Sprintf("http://api.weatherapi.com/v1/forecast.json?key=%s&q=%s&days=7&aqi=no&alerts=no", apiKey, city)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API request failed with status: %s", resp.Status)
	}

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
	fmt.Printf("\n📍 Current Weather in %s, %s\n", data.Location.Name, data.Location.Country)
	fmt.Printf("🌡️  Temperature: %.1f°C (Feels like: %.1f°C)\n", data.Current.TempC, data.Current.FeelsLikeC)
	fmt.Printf("☁️  Condition: %s\n", data.Current.Condition.Text)
	fmt.Printf("💧 Humidity: %.0f%%\n", data.Current.Humidity)
	fmt.Printf("💨 Wind Speed: %.1f km/h\n", data.Current.WindKph)
}

func (wa *WeatherAnalyzer) AnalyzeTemperatureTrends() {
	if len(wa.Data) == 0 || len(wa.Data[0].Forecast.ForecastDay) == 0 {
		fmt.Println("No forecast data available")
		return
	}

	forecastDays := wa.Data[0].Forecast.ForecastDay
	
	fmt.Printf("\n📈 7-Day Temperature Forecast Analysis\n")
	fmt.Println("=====================================")
	
	var totalMax, totalMin, totalAvg float64
	highestTemp := forecastDays[0].Day.MaxTempC
	lowestTemp := forecastDays[0].Day.MinTempC
	highestDay := forecastDays[0].Date
	lowestDay := forecastDays[0].Date

	for i, day := range forecastDays {
		totalMax += day.Day.MaxTempC
		totalMin += day.Day.MinTempC
		totalAvg += day.Day.AvgTempC

		if day.Day.MaxTempC > highestTemp {
			highestTemp = day.Day.MaxTempC
			highestDay = day.Date
		}
		if day.Day.MinTempC < lowestTemp {
			lowestTemp = day.Day.MinTempC
			lowestDay = day.Date
		}

		// Parse and format date
		parsedDate, _ := time.Parse("2006-01-02", day.Date)
		dayName := parsedDate.Format("Mon")
		
		fmt.Printf("%s: High %.1f°C / Low %.1f°C (Avg: %.1f°C)\n", 
			dayName, day.Day.MaxTempC, day.Day.MinTempC, day.Day.AvgTempC)
	}

	daysCount := float64(len(forecastDays))
	fmt.Printf("\n📊 Statistical Analysis:\n")
	fmt.Printf("Average High Temperature: %.1f°C\n", totalMax/daysCount)
	fmt.Printf("Average Low Temperature: %.1f°C\n", totalMin/daysCount)
	fmt.Printf("Overall Average Temperature: %.1f°C\n", totalAvg/daysCount)
	fmt.Printf("Highest Temperature: %.1f°C on %s\n", highestTemp, formatDate(highestDay))
	fmt.Printf("Lowest Temperature: %.1f°C on %s\n", lowestTemp, formatDate(lowestDay))
	fmt.Printf("Temperature Range: %.1f°C\n", highestTemp-lowestTemp)
}

func (wa *WeatherAnalyzer) GenerateVisualization() {
	if len(wa.Data) == 0 || len(wa.Data[0].Forecast.ForecastDay) == 0 {
		fmt.Println("No data available for visualization")
		return
	}

	forecastDays := wa.Data[0].Forecast.ForecastDay
	
	fmt.Printf("\n📊 Temperature Trend Visualization\n")
	fmt.Println("=================================")
	
	// Find min and max for scaling
	minTemp := forecastDays[0].Day.MinTempC
	maxTemp := forecastDays[0].Day.MaxTempC
	for _, day := range forecastDays {
		if day.Day.MinTempC < minTemp {
			minTemp = day.Day.MinTempC
		}
		if day.Day.MaxTempC > maxTemp {
			maxTemp = day.Day.MaxTempC
		}
	}

	// Create visualization for each day
	for i, day := range forecastDays {
		parsedDate, _ := time.Parse("2006-01-02", day.Date)
		dayName := parsedDate.Format("Mon")
		
		fmt.Printf("\n%s (%s):\n", dayName, day.Date)
		
		// Create temperature bar visualization
		createTemperatureBar(day.Day.MinTempC, day.Day.MaxTempC, minTemp, maxTemp)
		fmt.Printf("  Min: %.1f°C | Max: %.1f°C | Avg: %.1f°C\n", 
			day.Day.MinTempC, day.Day.MaxTempC, day.Day.AvgTempC)
	}
}

func createTemperatureBar(minTemp, maxTemp, overallMin, overallMax float64) {
	const barWidth = 50
	rangeSize := overallMax - overallMin
	
	// Calculate positions in the bar
	minPos := int(((minTemp - overallMin) / rangeSize) * barWidth)
	maxPos := int(((maxTemp - overallMin) / rangeSize) * barWidth)
	
	bar := make([]rune, barWidth+2)
	bar[0] = '|'
	bar[barWidth+1] = '|'
	
	for i := 1; i <= barWidth; i++ {
		if i >= minPos && i <= maxPos {
			bar[i] = '█' // Full block for temperature range
		} else {
			bar[i] = '─' // Dash for empty space
		}
	}
	
	fmt.Printf("  %s\n", string(bar))
}

func formatDate(dateStr string) string {
	parsedDate, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		return dateStr
	}
	return parsedDate.Format("Jan 02")
}