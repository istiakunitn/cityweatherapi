# Simple Application to Get Weather Data from OpenWeatherMap
This application fetches weather data from OpenWeatherMap for a given city.

## How to run
- Rename the `app.example.env` file to `app.env`.
- Ask a colleague for the required `environment variables` and paste them into the `app.env` file.
- Ensure that the `Docker daemon` is running on your PC.
- Run `docker compose build`.
- Run `docker compose up`.

The app will be available at `http://localhost:8080/`.

## Features

- Visiting `http://localhost:8080/` displays an HTML page where you can search for weather information by city name.
- After entering a city and clicking the search button, a card will appear with basic weather details.
- The search button fetches data from `/api/weather?city=<city_name>`.
- An `in-memory cache` stores recent `query results` for a city. If the same city is searched within five minutes, the data will be retrieved from the cache instead of making a new request.
- You can also send a `GET` request to `http://localhost:8080/api/weather?city=<city_name>` using `Postman` or `Thunder Client` to receive weather data in `JSON format`.

