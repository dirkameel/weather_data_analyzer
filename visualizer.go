package main

import (
	"fmt"
	"math"
	"strings"
)

func generateVisualization(data []WeatherData) {
	fmt.Println("\n=== TEMPERATURE VISUALIZATION ===")
	
	if len(data) == 0 {
		fmt.Println("No data to visualize")
		return
	}

	// Find temperature range for scaling
	minTemp, maxTemp := findMinMax(getTemperatures(data))
	tempRange := maxTemp - minTemp
	if tempRange == 0 {
		tempRange = 1 // Avoid division by zero
	}

	fmt.Println("\nTemperature Chart:")
	fmt.Println("=================")
	
	for _, wd := range data {
		// Scale temperature to bar length (0-50 characters)
		barLength := int(((wd.Temp - minTemp) / tempRange) * 50)
		if barLength < 1 {
			barLength = 1
		}
		
		bar := strings.Repeat("█", barLength)
		empty := strings.Repeat(" ", 50-barLength)
		
		fmt.Printf("%-15s |%s%s| %.1f°C\n", wd.City, bar, empty, wd.Temp)
	}

	// Add temperature scale
	fmt.Println("\nScale:")
	scaleStep := tempRange / 5
	for i := 0; i <= 5; i++ {
		temp := minTemp + (scaleStep * float64(i))
		fmt.Printf("%.1f°C ", temp)
	}
	fmt.Println()

	// Temperature comparison
	fmt.Println("\nTemperature Differences:")
	reference := data[0].Temp
	for i := 1; i < len(data); i++ {
		diff := data[i].Temp - reference
		comparison := "colder than"
		if diff > 0 {
			comparison = "warmer than"
		}
		fmt.Printf("%s is %.1f°C %s %s\n", 
			data[i].City, math.Abs(diff), comparison, data[0].City)
	}
}

func getTemperatures(data []WeatherData) []float64 {
	var temps []float64
	for _, wd := range data {
		temps = append(temps, wd.Temp)
	}
	return temps
}