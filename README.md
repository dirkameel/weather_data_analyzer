## Weather Data Analyzer

A simple Go application that fetches and analyzes weather data with temperature trend visualization.

### Features

- 🌡️ Current weather information
- 📈 3-day temperature forecasts
- 📊 ASCII-based temperature visualization
- 🔍 Temperature trend analysis
- 🌍 Multi-location support

### Prerequisites

1. Go 1.16 or higher
2. WeatherAPI.com account (free tier available)

### Setup

1. **Get API Key:**
   - Sign up at [WeatherAPI.com](https://www.weatherapi.com/)
   - Get your free API key

2. **Configure API Key:**
   ```bash
   # Option 1: Set environment variable
   export WEATHER_API_KEY="your_actual_api_key_here"
   
   # Option 2: Update config.go file
   # Change the default API key in config.go
   ```

3. **Install Dependencies:**
   ```bash
   go mod init weather-analyzer
   go mod tidy
   ```

### Usage

```bash
# Basic usage
go run main.go London

# Multiple cities
go run main.go "New York"
go run main.go Tokyo
go run main.go "Rio de Janeiro"
```

### Example Output

```
=== CURRENT WEATHER ===
Location: London, United Kingdom
Temperature: 15.5°C (Feels like: 14.2°C)
Condition: Partly cloudy
Humidity: 65%
Wind: 12.3 km/h

=== TEMPERATURE TRENDS (3-Day Forecast) ===
Monday (2024-01-15):
  Max: 16.2°C | Min: 8.5°C | Avg: 12.3°C
  Trend: warming (Δ1.2°C)

=== TEMPERATURE VISUALIZATION ===
Mon:  12.3°C [          ▀▀▀▀▀] (8.5°C - 16.2°C)
Tue:  13.5°C [           ▀▀▀▀▀▀] (9.1°C - 17.8°C)
Wed:  14.2°C [            ▀▀▀▀▀▀▀] (10.2°C - 18.5°C)
```

### Project Structure

- `main.go` - Main application logic and weather data fetching
- `config.go` - Configuration settings and API management
- `utils.go` - Utility functions for analysis and visualization
- `go.mod` - Go module definition

### API Notes

- Uses WeatherAPI.com (free tier: 1M calls/month)
- Provides current weather + 3-day forecast
- Data includes temperature, humidity, wind, and conditions

### Customization

You can modify:
- Number of forecast days in `config.go`
- Temperature units (metric/imperial)
- Visualization style in `utils.go`
- Add more weather parameters as needed

### Error Handling

- Handles API connection errors
- Validates city names
- Manages missing API keys
- Graceful degradation for visualization

This provides a solid foundation for weather data analysis that you can extend with additional features like historical data, multiple locations comparison, or more sophisticated visualizations.