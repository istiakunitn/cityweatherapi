package services

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/istiakunitn/cityweatherapi/config"
	"github.com/istiakunitn/cityweatherapi/interfaces"
)

func GetFromOpenWeather(city string) (interfaces.OpenWeatherResponse, error) {
	config, err := config.LoadConfig(".")

	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	url := fmt.Sprintf("%s?q=%s&appid=%s&units=metric", config.OpenWeatherBaseUrl, city, config.OpenWeatherApiKey)

	resp, err := http.Get(url)
	if err != nil {
		return interfaces.OpenWeatherResponse{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return interfaces.OpenWeatherResponse{}, fmt.Errorf("failed to get weather data: %s", resp.Status)
	}

	var weatherData interfaces.OpenWeatherResponse
	if err := json.NewDecoder(resp.Body).Decode(&weatherData); err != nil {
		return interfaces.OpenWeatherResponse{}, err
	}

	return weatherData, nil
}
