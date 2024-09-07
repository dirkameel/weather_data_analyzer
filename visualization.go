package main

import (
	"fmt"
	"math"
	"strings"
)

func generateVisualization(analysis *WeatherAnalysis) {
	fmt.Println("\n📊 TEMPERATURE VISUALIZATION")
	fmt.Println("============================")
	
	// Create a simple bar chart visualization
	temp := analysis.Temperature
	normalizedTemp := int(math.Abs(temp))
	
	if temp < 0 {
		fmt.Printf("Below Freezing: ")
		fmt.Println(strings.Repeat("❄️", normalizedTemp/2))
	} else {
		fmt.Printf("Temperature Scale: ")
		fmt.Println(strings.Repeat("🌡️", normalizedTemp/2))
	}
	
	// Temperature gauge
	fmt.Printf("\nTemperature Gauge:\n")
	fmt.Printf("-30°C ")
	printGauge(temp, -30, 40)
	fmt.Printf(" 40°C\n")
	
	// Trend indicator
	fmt.Printf("\nTrend Indicator: ")
	switch analysis.Trend {
	case "❄️ Freezing":
		fmt.Println("⬇️⬇️⬇️ (Extreme Cold)")
	case "🥶 Cold":
		fmt.Println("⬇️⬇️ (Cold)")
	case "😊 Mild":
		fmt.Println("➡️ (Moderate)")
	case "☀️ Warm":
		fmt.Println("⬆️⬆️ (Warm)")
	case "🔥 Hot":
		fmt.Println("⬆️⬆️⬆️ (Extreme Heat)")
	}
}

func printGauge(temp, min, max float64) {
	gaugeWidth := 20
	position := int((temp - min) / (max - min) * float64(gaugeWidth))
	
	for i := 0; i < gaugeWidth; i++ {
		if i == position {
			fmt.Print("📍")
		} else if i < position {
			fmt.Print("─")
		} else {
			fmt.Print("─")
		}
	}
}