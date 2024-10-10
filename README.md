# Weather Data Analyzer

A simple Go application that fetches, analyzes, and visualizes weather data from OpenWeatherMap API.

## Setup Instructions

1. **Get an API Key:**
   - Sign up at [OpenWeatherMap](https://openweathermap.org/api)
   - Get your free API key
   - Replace `your_api_key_here` in `api_client.go` with your actual API key

2. **Install Go:**
   - Make sure you have Go 1.21 or later installed
   - Verify with: `go version`

3. **Run the Application:**
   ```bash
   # Build and run directly
   go run main.go London Tokyo NewYork Paris Berlin
   
   # Or build an executable
   go build -o weather-analyzer
   ./weather-analyzer London Tokyo NewYork
   ```

## Usage Examples

```bash
# Compare multiple cities
go run main.go London Tokyo NewYork Sydney Mumbai

# Single city analysis
go run main.go Paris

# Regional comparison
go run main.go Seattle SanFrancisco LosAngeles Miami
```

## Features

- **Real-time Data**: Fetches current weather from OpenWeatherMap API
- **Temperature Analysis**: Ranks cities by temperature with descriptions
- **Visual Charts**: ASCII bar charts for easy temperature comparison
- **Data Persistence**: Saves weather data to `weather_data.json`
- **Humidity Tracking**: Includes humidity levels in analysis

## Output Includes

- Temperature rankings (hottest to coldest)
- Temperature descriptions (Freezing, Cold, Cool, Mild, Warm, Hot)
- ASCII bar chart visualization
- Temperature differences between cities
- Humidity percentages
- Overall statistics

## Requirements

- Go 1.21+
- Internet connection
- Valid OpenWeatherMap API key
- Cities must be spelled correctly as recognized by the API

## File Structure

- `main.go` - Application entry point and orchestration
- `api_client.go` - HTTP client for weather API
- `analyzer.go` - Data analysis and statistics
- `visualizer.go` - ASCII chart generation
- `go.mod` - Go module definition

The application will create a `weather_data.json` file containing the fetched weather data for future reference.