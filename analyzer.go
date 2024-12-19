package main

import (
	"fmt"
	"math"
	"strings"
)

func (wa *WeatherAnalyzer) DisplayCurrentWeather() {
	if len(wa.Data) == 0 {
		fmt.Println("No weather data available")
		return
	}

	data := wa.Data[0]
	
	fmt.Printf("\n📍 Current Weather in %s, %s\n", data.Location.Name, data.Location.Country)
	fmt.Println("====================================")
	fmt.Printf("🌡️  Temperature: %.1f°C (Feels like: %.1f°C)\n", data.Current.TempC, data.Current.FeelsLikeC)
	fmt.Printf("☁️  Condition: %s\n", data.Current.Condition.Text)
	fmt.Printf("💧 Humidity: %d%%\n", data.Current.Humidity)
	fmt.Printf("💨 Wind: %.1f km/h\n", data.Current.WindKph)
}

func (wa *WeatherAnalyzer) AnalyzeTemperatureTrends() {
	if len(wa.Data) == 0 {
		return
	}

	data := wa.Data[0]
	forecastDays := data.Forecast.Forecastday

	if len(forecastDays) == 0 {
		return
	}

	fmt.Printf("\n📈 7-Day Temperature Trends for %s\n", data.Location.Name)
	fmt.Println("====================================")

	// Calculate statistics
	var maxTemps, minTemps, avgTemps []float64
	var totalMax, totalMin, totalAvg float64

	for i, day := range forecastDays {
		maxTemps = append(maxTemps, day.Day.MaxTempC)
		minTemps = append(minTemps, day.Day.MinTempC)
		avgTemps = append(avgTemps, day.Day.AvgTempC)
		
		totalMax += day.Day.MaxTempC
		totalMin += day.Day.MinTempC
		totalAvg += day.Day.AvgTempC

		// Parse date for display
		date, _ := time.Parse("2006-01-02", day.Date)
		dayName := date.Format("Monday")

		fmt.Printf("%s: Max: %.1f°C, Min: %.1f°C, Avg: %.1f°C\n", 
			dayName, day.Day.MaxTempC, day.Day.MinTempC, day.Day.AvgTempC)
	}

	// Calculate averages
	days := float64(len(forecastDays))
	avgMax := totalMax / days
	avgMin := totalMin / days
	overallAvg := totalAvg / days

	// Find extremes
	maxTemp, minTemp := findExtremes(maxTemps, minTemps)
	tempRange := maxTemp - minTemp

	fmt.Printf("\n📊 Temperature Statistics:\n")
	fmt.Printf("Average High: %.1f°C\n", avgMax)
	fmt.Printf("Average Low: %.1f°C\n", avgMin)
	fmt.Printf("Overall Average: %.1f°C\n", overallAvg)
	fmt.Printf("Temperature Range: %.1f°C\n", tempRange)
	fmt.Printf("Highest Temp: %.1f°C\n", maxTemp)
	fmt.Printf("Lowest Temp: %.1f°C\n", minTemp)

	// Trend analysis
	trend := analyzeTrend(avgTemps)
	fmt.Printf("Trend: %s\n", trend)
}

func findExtremes(maxTemps, minTemps []float64) (float64, float64) {
	maxTemp := maxTemps[0]
	minTemp := minTemps[0]

	for _, temp := range maxTemps {
		if temp > maxTemp {
			maxTemp = temp
		}
	}

	for _, temp := range minTemps {
		if temp < minTemp {
			minTemp = temp
		}
	}

	return maxTemp, minTemp
}

func analyzeTrend(temps []float64) string {
	if len(temps) < 2 {
		return "Insufficient data for trend analysis"
	}

	// Simple linear regression for trend
	var sumX, sumY, sumXY, sumX2 float64
	n := float64(len(temps))

	for i, temp := range temps {
		x := float64(i)
		sumX += x
		sumY += temp
		sumXY += x * temp
		sumX2 += x * x
	}

	slope := (n*sumXY - sumX*sumY) / (n*sumX2 - sumX*sumX)

	if math.Abs(slope) < 0.1 {
		return "Stable"
	} else if slope > 0 {
		return "Warming trend"
	} else {
		return "Cooling trend"
	}
}

func (wa *WeatherAnalyzer) VisualizeTemperatureTrends() {
	if len(wa.Data) == 0 {
		return
	}

	data := wa.Data[0]
	forecastDays := data.Forecast.Forecastday

	if len(forecastDays) == 0 {
		return
	}

	fmt.Printf("\n📊 Temperature Visualization\n")
	fmt.Println("============================")

	// Find min and max for scaling
	minTemp, maxTemp := findExtremes(
		getMaxTemps(forecastDays),
		getMinTemps(forecastDays),
	)

	// Adjust range for better visualization
	rangeAdjust := (maxTemp - minTemp) * 0.1
	displayMin := minTemp - rangeAdjust
	displayMax := maxTemp + rangeAdjust
	tempRange := displayMax - displayMin

	const chartWidth = 50

	for i, day := range forecastDays {
		date, _ := time.Parse("2006-01-02", day.Date)
		dayName := date.Format("Mon")

		// Calculate bar positions
		maxPos := int(((day.Day.MaxTempC - displayMin) / tempRange) * chartWidth)
		minPos := int(((day.Day.MinTempC - displayMin) / tempRange) * chartWidth)
		avgPos := int(((day.Day.AvgTempC - displayMin) / tempRange) * chartWidth)

		// Create visualization bar
		bar := make([]rune, chartWidth+10)
		for j := range bar {
			if j < minPos {
				bar[j] = ' '
			} else if j == minPos {
				bar[j] = '❄' // Low temp marker
			} else if j == maxPos {
				bar[j] = '🔥' // High temp marker
			} else if j == avgPos {
				bar[j] = '●' // Average marker
			} else if j > minPos && j < maxPos {
				bar[j] = '─'
			} else {
				bar[j] = ' '
			}
		}

		fmt.Printf("%s: %s Max:%.1f°C\n", dayName, string(bar), day.Day.MaxTempC)
		fmt.Printf("      %s Min:%.1f°C\n", strings.Repeat(" ", minPos), day.Day.MinTempC)
		
		if i < len(forecastDays)-1 {
			fmt.Println()
		}
	}

	fmt.Printf("\nLegend: ❄ Low | ● Avg | 🔥 High\n")
}

func getMaxTemps(days []struct {
	Date string `json:"date"`
	Day  struct {
		MaxTempC float64 `json:"maxtemp_c"`
		MinTempC float64 `json:"mintemp_c"`
		AvgTempC float64 `json:"avgtemp_c"`
	} `json:"day"`
}) []float64 {
	var temps []float64
	for _, day := range days {
		temps = append(temps, day.Day.MaxTempC)
	}
	return temps
}

func getMinTemps(days []struct {
	Date string `json:"date"`
	Day  struct {
		MaxTempC float64 `json:"maxtemp_c"`
		MinTempC float64 `json:"mintemp_c"`
		AvgTempC float64 `json:"avgtemp_c"`
	} `json:"day"`
}) []float64 {
	var temps []float64
	for _, day := range days {
		temps = append(temps, day.Day.MinTempC)
	}
	return temps
}