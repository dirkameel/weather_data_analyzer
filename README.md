## Weather Data Analyzer

A simple Go application that fetches weather data from OpenWeatherMap API, stores it, analyzes temperature trends, and provides basic visualization.

### Prerequisites

1. Go 1.21 or later
2. OpenWeatherMap API key (free tier available at https://openweathermap.org/api)

### Setup Instructions

1. **Get API Key:**
   - Register at https://openweathermap.org/api
   - Get your free API key

2. **Configure the application:**
   - Edit `config.json` and replace `YOUR_API_KEY_HERE` with your actual API key
   - You can change the city name as needed

3. **Install dependencies:**
   ```bash
   go mod tidy
   ```

### Usage

1. **Run the application:**
   ```bash
   go run main.go weather_service.go storage.go analysis.go
   ```

2. **First run:**
   - The app will fetch current weather data
   - Create `weather_data.json` to store historical data
   - Display analysis and simple chart

3. **Subsequent runs:**
   - Each run adds new data point
   - Analysis includes trends from last 24 hours
   - Visualizes temperature changes over time

### Features

- **Data Collection:** Fetches real-time weather data from OpenWeatherMap API
- **Storage:** Maintains 24 hours of historical data in JSON format
- **Analysis:** 
  - Average, min, max temperatures
  - Temperature trends (warming/cooling/stable)
  - Weather recommendations based on temperature
- **Visualization:** Simple terminal-based chart showing temperature trends
- **Error Handling:** Robust error handling for API failures and file operations

### Output Example

```
Weather data processed successfully!
Current temperature in London: 15.5°C

=== WEATHER ANALYSIS ===
Data Points: 8
Time Period: 3.5 hours
Average Temperature: 16.2°C
Temperature Range: 2.5°C (Min: 15.0°C, Max: 17.5°C)
Trend: warming
Recommendation: Moderate temperature. Light jacket recommended.
========================

TEMPERATURE TREND CHART:
Time                | Temp (°C)
--------------------|-----------
10:30:15            |   15.0°C
11:00:22            |   15.5°C
11:30:18            |   16.0°C
12:00:45            |   16.5°C
12:30:33            |   17.0°C
13:00:29            |   17.5°C
13:30:12            |   17.0°C
14:00:08            |   16.5°C
```

### Files Description

- `main.go` - Application entry point and orchestration
- `config.json` - Configuration file for API key and city
- `weather_service.go` - API communication and data fetching
- `storage.go` - Data persistence and management
- `analysis.go` - Data analysis and visualization logic
- `go.mod` - Go module definition

### Customization

- Change city in `config.json`
- Modify data retention period in `storage.go` (currently 24 hours)
- Extend analysis features in `analysis.go`
- Add more weather parameters from the API response

### Error Handling

- Invalid API keys are handled gracefully
- Network failures are caught and reported
- File operations include proper error checking
- Missing configuration files are detected