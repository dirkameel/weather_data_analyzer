## Weather Data Analyzer

A simple Go application that fetches weather data from WeatherAPI and provides temperature trend analysis with ASCII visualization.

### Features

- 🌤️ Fetches current weather and 7-day forecast
- 📊 Analyzes temperature trends
- 📈 Generates ASCII temperature charts
- 📱 Simple command-line interface
- 🐳 Docker support

### Prerequisites

1. Go 1.21 or later
2. WeatherAPI account (free tier available)

### Setup

1. **Get API Key:**
   - Sign up at [WeatherAPI.com](https://www.weatherapi.com/)
   - Get your free API key

2. **Configure API Key (choose one method):**

   **Option A: Environment Variable**
   ```bash
   export WEATHER_API_KEY="your_api_key_here"
   ```

   **Option B: Config File**
   ```bash
   cp config.example.json config.json
   # Edit config.json with your API key
   ```

### Installation & Usage

**Method 1: Using Go Run**
```bash
go run main.go
# Or specify a city:
go run main.go "New York"
```

**Method 2: Build and Run**
```bash
make build
./weather-analyzer "London"
```

**Method 3: Docker**
```bash
# Build image
make docker-build

# Run with environment variable
docker run -e WEATHER_API_KEY=your_key_here weather-analyzer "Paris"
```

### Example Output

```
🌤️  Weather Data Analyzer
=========================
Enter city name (or press Enter for London): 

📍 Location: London, United Kingdom
🌡️  Current Temperature: 12.5°C

📊 7-Day Temperature Forecast:
Date     | Max Temp | Min Temp | Avg Temp | Trend
---------|----------|----------|----------|-------
Dec 15   |   14.2°C |    8.1°C |   11.2°C | ➡️
Dec 16   |   13.8°C |    7.9°C |   10.9°C | 📉
...

📊 Temperature Chart:
====================
 16.0°C |           
 15.0°C |           
 14.0°C |   █       
    ... |   █ ▄ █   
  8.0°C | █ █ █ █ █
```

### API Rate Limits

- Free tier: 1,000,000 calls per month
- No credit card required for free tier

### Project Structure

- `main.go` - Main application logic
- `config.example.json` - Example configuration file
- `go.mod` - Go module file
- `Dockerfile` - Container configuration
- `Makefile` - Build automation

### Customization

You can modify the chart appearance by adjusting the `chartHeight` variable in the `generateTemperatureChart` function.

### Error Handling

- Invalid API keys
- Network connectivity issues
- Invalid city names
- API rate limits

### License

MIT License - Feel free to modify and distribute.