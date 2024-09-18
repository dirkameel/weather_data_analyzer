.PHONY: build run test clean

build:
	@echo "Building weather analyzer..."
	go build -o weather-analyzer main.go visualization.go config.go

run:
	@echo "Running weather analyzer..."
	@if [ -z "$(city)" ]; then \
		echo "Usage: make run city=<city_name>"; \
		echo "Example: make run city=London"; \
	else \
		go run main.go visualization.go config.go $(city); \
	fi

test:
	@echo "Running tests..."
	go test -v ./...

clean:
	@echo "Cleaning up..."
	rm -f weather-analyzer
	rm -f weather.log

install-deps:
	@echo "Installing dependencies..."
	go mod tidy

setup:
	@echo "Setup instructions:"
	@echo "1. Get API key from https://www.weatherapi.com/"
	@echo "2. Set environment variable: export WEATHER_API_KEY=your_api_key_here"
	@echo "3. Run: make install-deps"
	@echo "4. Run: make run city=London"