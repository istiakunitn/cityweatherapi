# Simple application to get weather data from open weather map
This application is to get weather data from OpenWeathermap

## Dependency

`go 1.23.4` is used

## How to run?

- Clone the app 
- Add app.env file on project root with needed variables
- Run `go run server.go`

The app will listen on `localhost:8080`

## features

- If you visit on `http://localhost:8080/`, a html page is there to search weather by city name

- After entering a city and hitting search button a card will be shown with basic weather information

- Search button consumes data from `/api/weather?city=<city_name>`

- In memory cache is there to store recent query result for a city. If same city is search within five minutes, data will be fetched from in memory caching system

- You can also send a get request at `http://localhost:8080/api/weather?city=<city_name>` using postman or Thunder client to get weather json data

