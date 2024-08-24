package main

import (
	"fmt"
	"math"
	"strings"
)

// TemperatureCategory returns a descriptive category for the temperature
func TemperatureCategory(tempC float64) string {
	switch {
	case tempC < 0:
		return "Freezing ❄️"
	case tempC < 10:
		return "Cold 🥶"
	case tempC < 20:
		return "Cool 😊"
	case tempC < 30:
		return "Warm ☀️"
	default:
		return "Hot 🔥"
	}
}

// GetTrendIndicator returns arrow indicators for temperature trends
func GetTrendIndicator(current, previous float64) string {
	diff := current - previous
	switch {
	case diff > 2:
		return "↗️"
	case diff > 0.5:
		return "↗️"
	case diff < -2:
		return "↘️"
	case diff < -0.5:
		return "↘️"
	default:
		return "→"
	}
}

// CreateSparkline creates a simple sparkline for temperature visualization
func CreateSparkline(temps []float64, width int) string {
	if len(temps) == 0 {
		return ""
	}

	min := temps[0]
	max := temps[0]
	for _, temp := range temps {
		if temp < min {
			min = temp
		}
		if temp > max {
			max = temp
		}
	}

	rangeVal := max - min
	if rangeVal == 0 {
		rangeVal = 1 // Avoid division by zero
	}

	sparks := []string{"▁", "▂", "▃", "▄", "▅", "▆", "▇", "█"}
	sparkline := make([]string, len(temps))

	for i, temp := range temps {
		pos := int(((temp - min) / rangeVal) * float64(len(sparks)-1))
		if pos < 0 {
			pos = 0
		}
		if pos >= len(sparks) {
			pos = len(sparks) - 1
		}
		sparkline[i] = sparks[pos]
	}

	return strings.Join(sparkline, "")
}

// Round rounds a float64 to specified decimal places
func Round(value float64, places int) float64 {
	shift := math.Pow(10, float64(places))
	return math.Round(value*shift) / shift
}

// FormatTemp nicely formats temperature with degree symbol
func FormatTemp(temp float64) string {
	return fmt.Sprintf("%.1f°C", temp)
}