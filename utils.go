package main

import (
	"fmt"
	"math"
	"strings"
)

// TemperatureAnalysis provides detailed temperature insights
type TemperatureAnalysis struct {
	CurrentTemp    float64
	AverageTemp    float64
	MaxTemp        float64
	MinTemp        float64
	TemperatureRange float64
	Trend          string
}

func AnalyzeTemperatures(data WeatherData) TemperatureAnalysis {
	var maxTemp, minTemp, sumTemp float64
	days := data.Forecast.Forecastday
	
	if len(days) == 0 {
		return TemperatureAnalysis{}
	}
	
	minTemp = days[0].Day.MintempC
	maxTemp = days[0].Day.MaxtempC
	
	for _, day := range days {
		sumTemp += day.Day.AvgtempC
		if day.Day.MintempC < minTemp {
			minTemp = day.Day.MintempC
		}
		if day.Day.MaxtempC > maxTemp {
			maxTemp = day.Day.MaxtempC
		}
	}
	
	avgTemp := sumTemp / float64(len(days))
	
	// Determine trend
	trend := "stable"
	if len(days) > 1 {
		firstAvg := days[0].Day.AvgtempC
		lastAvg := days[len(days)-1].Day.AvgtempC
		change := lastAvg - firstAvg
		
		if math.Abs(change) > 2 {
			if change > 0 {
				trend = "warming"
			} else {
				trend = "cooling"
			}
		}
	}
	
	return TemperatureAnalysis{
		CurrentTemp:     data.Current.TempC,
		AverageTemp:     avgTemp,
		MaxTemp:         maxTemp,
		MinTemp:         minTemp,
		TemperatureRange: maxTemp - minTemp,
		Trend:           trend,
	}
}

// CreateASCIIChart creates a simple ASCII bar chart for temperatures
func CreateASCIIChart(days []struct {
	Date string
	Day  struct {
		MaxtempC float64
		MintempC float64
		AvgtempC float64
	}
}) string {
	var chart strings.Builder
	chart.WriteString("\n📊 Temperature Chart:\n")
	chart.WriteString("    Min  ─── Avg ─── Max\n")
	
	for i, day := range days {
		// Scale temperatures for visualization (assuming -10°C to 40°C range)
		minPos := scaleTemperature(day.Day.MintempC, -10, 40, 20)
		avgPos := scaleTemperature(day.Day.AvgtempC, -10, 40, 20)
		maxPos := scaleTemperature(day.Day.MaxtempC, -10, 40, 20)
		
		chart.WriteString(fmt.Sprintf("Day %d: ", i+1))
		
		// Create the bar
		for pos := 0; pos <= 20; pos++ {
			if pos == minPos {
				chart.WriteString("L")
			} else if pos == maxPos {
				chart.WriteString("H")
			} else if pos == avgPos {
				chart.WriteString("●")
			} else if pos > minPos && pos < maxPos {
				chart.WriteString("─")
			} else {
				chart.WriteString(" ")
			}
		}
		
		chart.WriteString(fmt.Sprintf(" %.1f°C\n", day.Day.AvgtempC))
	}
	
	return chart.String()
}

func scaleTemperature(temp, minRange, maxRange float64, width int) int {
	scaled := (temp - minRange) / (maxRange - minRange) * float64(width)
	return int(math.Round(scaled))
}

// GetWeatherEmoji returns an emoji based on temperature
func GetWeatherEmoji(temp float64) string {
	switch {
	case temp < 0:
		return "❄️"
	case temp < 10:
		return "☁️"
	case temp < 20:
		return "⛅"
	case temp < 30:
		return "☀️"
	default:
		return "🔥"
	}
}