# Simple application to get weather data from open weather map
This application is to get weather data from OpenWeathermap

## Dependency

`go 1.23.4` is used

## How to run?

Clone the app and just run `go run server.go`

The app will listen on `localhost:8080`

You can send a get request in endpoint `/api/weather?city=<city_name>` and it will return weather data json

