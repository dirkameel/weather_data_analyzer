# 🌤️ Go Weather Data Analyzer

A simple Go application that fetches weather data from WeatherAPI.com and provides temperature trend analysis with visualization.

## 🚀 Features

- **Real-time Weather Data**: Fetches current weather and 3-day forecast
- **Temperature Analysis**: Provides daily and hourly temperature trends
- **ASCII Visualization**: Creates simple temperature charts in terminal
- **Trend Analysis**: Identifies warmest/coldest times and temperature ranges

## 📋 Prerequisites

1. **Go 1.21+** installed on your system
2. **WeatherAPI.com account** for free API key

## 🛠️ Setup Instructions

### 1. Get API Key
1. Visit [WeatherAPI.com](https://www.weatherapi.com/)
2. Sign up for a free account
3. Get your API key from the dashboard

### 2. Configure the Application
1. Copy `config.example.go` to `config.go`:
   ```bash
   cp config.example.go config.go
   ```
2. Edit `config.go` and replace `YOUR_API_KEY_HERE` with your actual API key

### 3. Install Dependencies
```bash
go mod tidy
```

## 🎯 Usage

### Basic Usage
```bash
go run main.go London
```

### Examples
```bash
# Single word location
go run main.go Tokyo

# Multi-word location
go run main.go "New York"

# With country code
go run main.go "Paris, France"
```

## 📊 Output Example

```
🌤️  Fetching weather data for: London

📍 Current Weather in London, United Kingdom
🌡️  Temperature: 15.5°C (59.9°F)
==================================================

📊 Temperature Trends Analysis
==================================================

Today (2024-01-15):
  📈 High: 16.2°C
  📉 Low: 8.5°C
  📊 Average: 12.3°C
  📏 Range: 7.7°C

📈 Temperature Visualization
==================================================

Today's Temperature Chart:
Time  | Temp°C | Chart
------|--------|-------------------
00:00 |   10.2 | █████
03:00 |    8.5 | ████
06:00 |    9.1 | ████
...
```

## 📁 Project Structure

- `main.go` - Entry point and orchestration
- `api_client.go` - HTTP client for Weather API
- `analyzer.go` - Weather data analysis and visualization
- `go.mod` - Go module definition
- `config.example.go` - API configuration template

## 🔧 API Limitations

- Free tier: 1,000,000 calls per month
- Rate limiting: No strict limits, but be reasonable
- Data refresh: Real-time with 15-minute cache

## 🐛 Troubleshooting

**API Key Issues:**
- Ensure your API key is correctly set in `config.go`
- Verify your API key is active in WeatherAPI dashboard

**Location Not Found:**
- Try different location formats (city, city+country, coordinates)
- Check spelling and use English location names

**Network Issues:**
- Verify internet connection
- Check if WeatherAPI.com is accessible

## 📝 License

This project is for educational purposes. Weather data provided by [WeatherAPI.com](https://www.weatherapi.com/).

## 🔮 Future Enhancements

- Add historical data analysis
- Support for multiple weather providers
- Export data to CSV/JSON
- Web interface version
- Mobile notifications for significant temperature changes

Enjoy analyzing weather trends! 🌈