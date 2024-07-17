package main

import (
	"fmt"
	"math"
	"time"
)

type AnalysisResult struct {
	AverageTemp    float64   `json:"average_temp"`
	MaxTemp        float64   `json:"max_temp"`
	MinTemp        float64   `json:"min_temp"`
	TempRange      float64   `json:"temp_range"`
	DataPoints     int       `json:"data_points"`
	TimePeriod     string    `json:"time_period"`
	Trend          string    `json:"trend"`
	Recommendation string    `json:"recommendation"`
}

func analyzeAndVisualize() error {
	data, err := loadWeatherData()
	if err != nil {
		return fmt.Errorf("could not load weather data: %v", err)
	}

	if len(data) == 0 {
		fmt.Println("No weather data available for analysis")
		return nil
	}

	result := analyzeData(data)
	displayAnalysis(result)
	displaySimpleChart(data)

	return nil
}

func analyzeData(data []WeatherData) AnalysisResult {
	if len(data) == 0 {
		return AnalysisResult{}
	}

	var sumTemp float64
	maxTemp := data[0].Temp
	minTemp := data[0].Temp

	for _, item := range data {
		sumTemp += item.Temp
		if item.Temp > maxTemp {
			maxTemp = item.Temp
		}
		if item.Temp < minTemp {
			minTemp = item.Temp
		}
	}

	avgTemp := sumTemp / float64(len(data))
	
	// Determine trend
	var trend string
	if len(data) >= 2 {
		latest := data[len(data)-1].Temp
		earliest := data[0].Temp
		if latest > earliest+0.5 {
			trend = "warming"
		} else if latest < earliest-0.5 {
			trend = "cooling"
		} else {
			trend = "stable"
		}
	} else {
		trend = "insufficient data"
	}

	// Generate recommendation
	recommendation := generateRecommendation(avgTemp)

	return AnalysisResult{
		AverageTemp:    math.Round(avgTemp*10) / 10,
		MaxTemp:        math.Round(maxTemp*10) / 10,
		MinTemp:        math.Round(minTemp*10) / 10,
		TempRange:      math.Round((maxTemp-minTemp)*10) / 10,
		DataPoints:     len(data),
		TimePeriod:     calculateTimePeriod(data),
		Trend:          trend,
		Recommendation: recommendation,
	}
}

func calculateTimePeriod(data []WeatherData) string {
	if len(data) < 2 {
		return "Single reading"
	}
	
	start := data[0].Timestamp
	end := data[len(data)-1].Timestamp
	duration := end.Sub(start)
	
	return fmt.Sprintf("%.1f hours", duration.Hours())
}

func generateRecommendation(avgTemp float64) string {
	switch {
	case avgTemp < 0:
		return "Very cold! Dress warmly with multiple layers."
	case avgTemp < 10:
		return "Cold weather. Wear a jacket and warm clothing."
	case avgTemp < 20:
		return "Moderate temperature. Light jacket recommended."
	case avgTemp < 30:
		return "Warm weather. Light clothing is comfortable."
	default:
		return "Hot weather. Stay hydrated and wear light clothes."
	}
}

func displayAnalysis(result AnalysisResult) {
	fmt.Println("\n=== WEATHER ANALYSIS ===")
	fmt.Printf("Data Points: %d\n", result.DataPoints)
	fmt.Printf("Time Period: %s\n", result.TimePeriod)
	fmt.Printf("Average Temperature: %.1f°C\n", result.AverageTemp)
	fmt.Printf("Temperature Range: %.1f°C (Min: %.1f°C, Max: %.1f°C)\n", 
		result.TempRange, result.MinTemp, result.MaxTemp)
	fmt.Printf("Trend: %s\n", result.Trend)
	fmt.Printf("Recommendation: %s\n", result.Recommendation)
	fmt.Println("========================\n")
}

func displaySimpleChart(data []WeatherData) {
	if len(data) == 0 {
		return
	}

	fmt.Println("TEMPERATURE TREND CHART:")
	fmt.Println("Time                | Temp (°C)")
	fmt.Println("--------------------|-----------")
	
	for i, item := range data {
		if i >= 10 { // Limit display to last 10 readings
			break
		}
		timeStr := item.Timestamp.Format("15:04:05")
		fmt.Printf("%-19s | %6.1f°C\n", timeStr, item.Temp)
	}
	fmt.Println()
}