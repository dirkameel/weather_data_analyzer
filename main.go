package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

// WeatherData represents the structure for weather API response
type WeatherData struct {
	Location struct {
		Name    string `json:"name"`
		Country string `json:"country"`
	} `json:"location"`
	Current struct {
		TempC     float64 `json:"temp_c"`
		TempF     float64 `json:"temp_f"`
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

// WeatherTrend represents temperature trends over time
type WeatherTrend struct {
	Location    string
	Dates       []string
	MaxTemps    []float64
	MinTemps    []float64
	AvgTemps    []float64
	CurrentTemp float64
}

// Config holds API configuration
type Config struct {
	APIKey string `json:"api_key"`
}

func main() {
	fmt.Println("🌤️  Weather Data Analyzer")
	fmt.Println("=========================")

	// Load configuration
	config, err := loadConfig()
	if err != nil {
		log.Fatal("Error loading config:", err)
	}

	// Get city from user input or use default
	city := getUserInput()
	
	// Fetch weather data
	weatherData, err := fetchWeatherData(city, config.APIKey)
	if err != nil {
		log.Fatal("Error fetching weather data:", err)
	}

	// Analyze and display trends
	trends := analyzeWeatherTrends(weatherData)
	displayWeatherAnalysis(trends)
	
	// Generate visualization
	generateTemperatureChart(trends)
	
	fmt.Println("\n✅ Analysis complete! Check the generated chart above.")
}

func loadConfig() (*Config, error) {
	// Try to load from config file first
	if _, err := os.Stat("config.json"); err == nil {
		file, err := os.ReadFile("config.json")
		if err != nil {
			return nil, err
		}
		
		var config Config
		err = json.Unmarshal(file, &config)
		if err != nil {
			return nil, err
		}
		
		if config.APIKey != "" {
			return &config, nil
		}
	}
	
	// Fallback to environment variable
	apiKey := os.Getenv("WEATHER_API_KEY")
	if apiKey == "" {
		return nil, fmt.Errorf("no API key found. Please set WEATHER_API_KEY environment variable or create config.json")
	}
	
	return &Config{APIKey: apiKey}, nil
}

func getUserInput() string {
	if len(os.Args) > 1 {
		return os.Args[1]
	}
	
	fmt.Print("Enter city name (or press Enter for London): ")
	var input string
	fmt.Scanln(&input)
	
	if input == "" {
		return "London"
	}
	return input
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
	
	body, err := io.ReadAll(resp.Body)
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

func analyzeWeatherTrends(data *WeatherData) *WeatherTrend {
	trend := &WeatherTrend{
		Location:    fmt.Sprintf("%s, %s", data.Location.Name, data.Location.Country),
		CurrentTemp: data.Current.TempC,
	}
	
	for _, day := range data.Forecast.Forecastday {
		trend.Dates = append(trend.Dates, formatDate(day.Date))
		trend.MaxTemps = append(trend.MaxTemps, day.Day.MaxTempC)
		trend.MinTemps = append(trend.MinTemps, day.Day.MinTempC)
		trend.AvgTemps = append(trend.AvgTemps, day.Day.AvgTempC)
	}
	
	return trend
}

func formatDate(dateStr string) string {
	t, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		return dateStr
	}
	return t.Format("Jan 02")
}

func displayWeatherAnalysis(trends *WeatherTrend) {
	fmt.Printf("\n📍 Location: %s\n", trends.Location)
	fmt.Printf("🌡️  Current Temperature: %.1f°C\n", trends.CurrentTemp)
	
	fmt.Println("\n📊 7-Day Temperature Forecast:")
	fmt.Println("Date     | Max Temp | Min Temp | Avg Temp | Trend")
	fmt.Println("---------|----------|----------|----------|-------")
	
	for i := 0; i < len(trends.Dates); i++ {
		trendSymbol := getTrendSymbol(trends.AvgTemps, i)
		fmt.Printf("%-8s | %7.1f°C | %7.1f°C | %7.1f°C | %s\n", 
			trends.Dates[i], trends.MaxTemps[i], trends.MinTemps[i], trends.AvgTemps[i], trendSymbol)
	}
	
	// Calculate statistics
	avgMax := calculateAverage(trends.MaxTemps)
	avgMin := calculateAverage(trends.MinTemps)
	tempRange := trends.MaxTemps[0] - trends.MinTemps[0]
	
	fmt.Printf("\n📈 Statistics:")
	fmt.Printf("\n• Average High: %.1f°C", avgMax)
	fmt.Printf("\n• Average Low: %.1f°C", avgMin)
	fmt.Printf("\n• Temperature Range: %.1f°C", tempRange)
	
	overallTrend := getOverallTrend(trends.AvgTemps)
	fmt.Printf("\n• Overall Trend: %s\n", overallTrend)
}

func getTrendSymbol(temps []float64, index int) string {
	if index == 0 || index >= len(temps)-1 {
		return "➡️" // stable
	}
	
	if temps[index] > temps[index-1] {
		return "📈" // rising
	} else if temps[index] < temps[index-1] {
		return "📉" // falling
	}
	return "➡️" // stable
}

func getOverallTrend(temps []float64) string {
	if len(temps) < 2 {
		return "Insufficient data"
	}
	
	firstAvg := calculateAverage(temps[:len(temps)/2])
	secondAvg := calculateAverage(temps[len(temps)/2:])
	
	if secondAvg > firstAvg + 1.0 {
		return "Warming trend 🔥"
	} else if secondAvg < firstAvg - 1.0 {
		return "Cooling trend ❄️"
	}
	return "Stable conditions ⚖️"
}

func calculateAverage(temps []float64) float64 {
	if len(temps) == 0 {
		return 0
	}
	
	sum := 0.0
	for _, temp := range temps {
		sum += temp
	}
	return sum / float64(len(temps))
}

func generateTemperatureChart(trends *WeatherTrend) {
	fmt.Println("\n📊 Temperature Chart:")
	fmt.Println("====================")
	
	// Find min and max for scaling
	minTemp, maxTemp := findMinMax(trends.MinTemps, trends.MaxTemps)
	chartHeight := 10
	
	for row := chartHeight; row >= 0; row-- {
		temp := minTemp + (float64(row) / float64(chartHeight)) * (maxTemp - minTemp)
		
		fmt.Printf("%5.1f°C | ", temp)
		
		for i := 0; i < len(trends.Dates); i++ {
			if trends.MaxTemps[i] >= temp && trends.MinTemps[i] <= temp {
				// Temperature falls within this day's range
				if trends.AvgTemps[i] >= temp {
					fmt.Print("█") // Above average
				} else {
					fmt.Print("▄") // Below average but in range
				}
			} else if trends.MaxTemps[i] >= temp {
				fmt.Print(" ") // Above max temp
			} else {
				fmt.Print(" ") // Below min temp
			}
			fmt.Print(" ")
		}
		fmt.Println()
	}
	
	// Print X-axis labels
	fmt.Print("       | ")
	for i, date := range trends.Dates {
		if i == 0 {
			fmt.Print("T")
		} else {
			fmt.Print(" ")
		}
		fmt.Print(" ")
	}
	fmt.Println()
	
	fmt.Print("       | ")
	for _, date := range trends.Dates {
		fmt.Printf("%-2s", string(date[0]))
	}
	fmt.Println()
	
	fmt.Print("       | ")
	for i := range trends.Dates {
		if i == len(trends.Dates)-1 {
			fmt.Print("→")
		} else {
			fmt.Print("  ")
		}
	}
	fmt.Println(" Time")
	
	fmt.Println("\nLegend: █ = Average/Hot, ▄ = Cool, empty = Outside range")
}

func findMinMax(minTemps, maxTemps []float64) (float64, float64) {
	min := minTemps[0]
	max := maxTemps[0]
	
	for i := 1; i < len(minTemps); i++ {
		if minTemps[i] < min {
			min = minTemps[i]
		}
		if maxTemps[i] > max {
			max = maxTemps[i]
		}
	}
	
	// Add some padding
	min = min - 2
	max = max + 2
	
	return min, max
}