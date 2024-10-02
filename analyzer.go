package main

import (
	"fmt"
	"sort"
)

type AnalysisResult struct {
	City        string
	AvgTemp     float64
	MaxTemp     float64
	MinTemp     float64
	Humidity    int
	Temperature string
}

func analyzeWeatherData(data []WeatherData) {
	fmt.Println("\n=== WEATHER ANALYSIS ===")
	
	if len(data) == 0 {
		fmt.Println("No data to analyze")
		return
	}

	var results []AnalysisResult
	var allTemps []float64

	for _, wd := range data {
		allTemps = append(allTemps, wd.Temp)
		
		tempDesc := getTemperatureDescription(wd.Temp)
		result := AnalysisResult{
			City:        wd.City,
			AvgTemp:     wd.Temp,
			MaxTemp:     wd.Temp,
			MinTemp:     wd.Temp,
			Humidity:    wd.Humidity,
			Temperature: tempDesc,
		}
		results = append(results, result)
	}

	// Sort by temperature
	sort.Slice(results, func(i, j int) bool {
		return results[i].AvgTemp > results[j].AvgTemp
	})

	// Display results
	fmt.Printf("\nTemperature Ranking (Hottest to Coldest):\n")
	fmt.Println("----------------------------------------")
	for i, result := range results {
		fmt.Printf("%d. %s: %.1f°C (%s) - Humidity: %d%%\n", 
			i+1, result.City, result.AvgTemp, result.Temperature, result.Humidity)
	}

	// Overall statistics
	if len(allTemps) > 0 {
		avg := calculateAverage(allTemps)
		min, max := findMinMax(allTemps)
		fmt.Printf("\nOverall Statistics:\n")
		fmt.Printf("Average Temperature: %.1f°C\n", avg)
		fmt.Printf("Temperature Range: %.1f°C to %.1f°C\n", min, max)
	}
}

func getTemperatureDescription(temp float64) string {
	switch {
	case temp < 0:
		return "Freezing"
	case temp < 10:
		return "Cold"
	case temp < 20:
		return "Cool"
	case temp < 25:
		return "Mild"
	case temp < 30:
		return "Warm"
	default:
		return "Hot"
	}
}

func calculateAverage(temps []float64) float64 {
	sum := 0.0
	for _, temp := range temps {
		sum += temp
	}
	return sum / float64(len(temps))
}

func findMinMax(temps []float64) (min, max float64) {
	if len(temps) == 0 {
		return 0, 0
	}
	min, max = temps[0], temps[0]
	for _, temp := range temps {
		if temp < min {
			min = temp
		}
		if temp > max {
			max = temp
		}
	}
	return min, max
}