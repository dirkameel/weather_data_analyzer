package main

import (
	"fmt"
	"math"
	"time"
)

func NewWeatherAnalyzer(data *WeatherData, location string) *WeatherAnalyzer {
	return &WeatherAnalyzer{
		data:     data,
		location: location,
	}
}

func (wa *WeatherAnalyzer) DisplayCurrentWeather() {
	fmt.Printf("\n📍 Current Weather in %s, %s\n", wa.data.Location.Name, wa.data.Location.Country)
	fmt.Printf("🌡️  Temperature: %.1f°C (%.1f°F)\n", wa.data.Current.TempC, wa.data.Current.TempF)
	fmt.Println("=" * 50)
}

func (wa *WeatherAnalyzer) AnalyzeTrends() {
	fmt.Printf("\n📊 Temperature Trends Analysis\n")
	fmt.Println("=" * 50)

	// Analyze daily trends
	for i, day := range wa.data.Forecast.Forecastday {
		dayName := "Today"
		if i == 1 {
			dayName = "Tomorrow"
		} else if i > 1 {
			dayName = fmt.Sprintf("Day %d", i+1)
		}

		fmt.Printf("\n%s (%s):\n", dayName, day.Date)
		fmt.Printf("  📈 High: %.1f°C\n", day.Day.MaxTempC)
		fmt.Printf("  📉 Low: %.1f°C\n", day.Day.MinTempC)
		fmt.Printf("  📊 Average: %.1f°C\n", day.Day.AvgTempC)
		fmt.Printf("  📏 Range: %.1f°C\n", day.Day.MaxTempC-day.Day.MinTempC)
	}

	// Analyze hourly trends for today
	todayHours := wa.data.Forecast.Forecastday[0].Hour
	wa.analyzeHourlyTrends(todayHours)
}

func (wa *WeatherAnalyzer) analyzeHourlyTrends(hours []HourData) {
	fmt.Printf("\n⏰ Today's Hourly Analysis:\n")
	
	var minTemp, maxTemp float64 = math.MaxFloat64, -math.MaxFloat64
	var warmestHour, coldestHour string
	var totalTemp float64

	for _, hour := range hours {
		if hour.TempC < minTemp {
			minTemp = hour.TempC
			coldestHour = formatHour(hour.TimeEpoch)
		}
		if hour.TempC > maxTemp {
			maxTemp = hour.TempC
			warmestHour = formatHour(hour.TimeEpoch)
		}
		totalTemp += hour.TempC
	}

	fmt.Printf("  🔥 Warmest time: %s (%.1f°C)\n", warmestHour, maxTemp)
	fmt.Printf("  ❄️  Coldest time: %s (%.1f°C)\n", coldestHour, minTemp)
	fmt.Printf("  📊 Daily average: %.1f°C\n", totalTemp/float64(len(hours)))
}

func (wa *WeatherAnalyzer) GenerateVisualization() {
	fmt.Printf("\n📈 Temperature Visualization\n")
	fmt.Println("=" * 50)
	
	// Simple ASCII visualization for today's hourly temperatures
	todayHours := wa.data.Forecast.Forecastday[0].Hour
	
	fmt.Printf("\nToday's Temperature Chart:\n")
	fmt.Println("Time  | Temp°C | Chart")
	fmt.Println("------|--------|-------------------")
	
	for i := 0; i < len(todayHours); i += 3 { // Show every 3 hours for brevity
		hour := todayHours[i]
		temp := hour.TempC
		timeStr := formatHour(hour.TimeEpoch)
		
		// Create a simple bar chart
		bars := int((temp + 10) / 2) // Scale for visualization
		if bars < 0 {
			bars = 0
		}
		if bars > 20 {
			bars = 20
		}
		
		chart := strings.Repeat("█", bars)
		fmt.Printf("%s | %6.1f | %s\n", timeStr, temp, chart)
	}
}

func formatHour(epoch int64) string {
	t := time.Unix(epoch, 0)
	return t.Format("15:04")
}