package main

const (
	// WeatherAPI configuration
	WeatherAPIBaseURL = "http://api.weatherapi.com/v1"
	DefaultDays       = 3
	
	// Visualization settings
	MaxChartWidth    = 50
	TempScaleOffset  = 10 // For handling negative temperatures in visualization
)

type Config struct {
	APIKey     string
	Days       int
	Units      string // "metric" or "imperial"
	EnableCache bool
}

func NewConfig() *Config {
	return &Config{
		APIKey:     getAPIKey(),
		Days:       DefaultDays,
		Units:      "metric",
		EnableCache: true,
	}
}

func getAPIKey() string {
	// First try environment variable
	if apiKey := os.Getenv("WEATHER_API_KEY"); apiKey != "" {
		return apiKey
	}
	
	// You can also read from a config file here
	return "your_api_key_here" // Default fallback
}