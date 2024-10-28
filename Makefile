.PHONY: build run docker-build docker-run clean

build:
	go build -o weather-analyzer

run: build
	./weather-analyzer

docker-build:
	docker build -t weather-analyzer .

docker-run:
	docker run -e WEATHER_API_KEY=your_api_key_here weather-analyzer

clean:
	rm -f weather-analyzer