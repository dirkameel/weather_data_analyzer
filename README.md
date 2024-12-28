# 🌤️ Weather Data Analyzer

A simple Go application that fetches and analyzes weather data from WeatherAPI with temperature trends visualization.

## Features

- 📍 **Current weather display** with temperature, humidity, wind, and conditions
- 📈 **7-day forecast** with detailed temperature analysis
- 📊 **Statistical analysis** including averages, extremes, and trends
- 🎯 **Visual temperature trends** with ASCII chart visualization
- ⚙️ **Configurable** with support for multiple locations

## Prerequisites

- Go 1.21 or later
- WeatherAPI account (free tier available)

## Setup

1. **Get API Key**:
   - Sign up at [WeatherAPI.com](https://www.weatherapi.com/)
   - Get your free API key

2. **Configure API Key**:
   - Option 1: Set environment variable:
     ```bash
     export WEATHER_API_KEY="your_api_key_here"
     ```
   - Option 2: Edit `config.json`:
     ```json
     {
       "api_key": "your_actual_api_key_here",
       "units": "metric",
       "default_city": "London"
     }
     ```

3. **Install Dependencies**:
   ```bash
   go mod tidy
   ```

## Usage

### Basic Usage
```bash
go run main.go analyzer.go config.go
```

### Specify Location
```bash
# As command line argument
go run . "New York"

# Or as interactive input
go run .
# Then enter location when prompted
```

### Examples
```bash
# Different cities
go run . "Tokyo"
go run . "Paris"
go run . "Sydney"

# Cities with spaces
go run . "New York"
go run . "Mexico City"
```

## Output Example

```
🌤️  Weather Data Analyzer
==========================
Enter location (or press Enter for London): Tokyo

📍 Current Weather in Tokyo, Japan
====================================
🌡️  Temperature: 15.2°C (Feels like: 14.8°C)
☁️  Condition: Partly cloudy
💧 Humidity: 65%
💨 Wind: 12.3 km/h

📈 7-Day Temperature Trends for Tokyo
====================================
Monday: Max: 16.5°C, Min: 8.2°C, Avg: 12.3°C
Tuesday: Max: 17.1°C, Min: 9.1°C, Avg: 13.1°C
...

📊 Temperature Statistics:
Average High: 16.8°C
Average Low: 9.5°C
Overall Average: 13.2°C
Temperature Range: 8.9°C
Highest Temp: 18.3°C
Lowest Temp: 7.4°C
Trend: Warming trend

📊 Temperature Visualization
============================
Mon: ─────────────❄─────────────────────────────●────────────🔥 Max:16.5°C
                Min:8.2°C
...
```

## File Structure

```
weather-analyzer/
├── main.go          # Main application entry point
├── analyzer.go      # Weather analysis and visualization logic
├── config.go        # Configuration management
├── config.json      # API configuration file
└── go.mod          # Go module definition
```

## API Rate Limits

- Free tier: 1,000,000 calls per month
- 1 call per location per execution
- Suitable for personal use and testing

## Error Handling

- Invalid API keys
- Network connectivity issues
- Invalid location names
- API rate limiting

## Customization

You can modify:
- Forecast days in `main.go` (change `&days=7`)
- Temperature units in `config.json`
- Visualization width in `analyzer.go`

## License

MIT License - Feel free to modify and distribute.