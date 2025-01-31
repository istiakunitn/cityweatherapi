package services

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/istiakunitn/cityweatherapi/config"
	"github.com/istiakunitn/cityweatherapi/interfaces"
)

const cacheExpiry = 5 * time.Minute

var (
	cache = make(map[string]interfaces.WeatherCacheEntry)
	mutex sync.Mutex
)

func GetFromOpenWeather(city string) (interfaces.OpenWeatherResponse, error) {
	fmt.Println("Checking start if data is fetched from cache")

	mutex.Lock()
	entry, found := cache[city]
	if found && time.Now().Before(entry.ExpiresAt) {
		mutex.Unlock()
		return entry.Data, nil
	}
	mutex.Unlock()

	config, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	fmt.Println("If you see me, you are calling open weather api to get weather data")

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

	mutex.Lock()
	cache[city] = interfaces.WeatherCacheEntry{Data: weatherData, ExpiresAt: time.Now().Add(cacheExpiry)}
	mutex.Unlock()

	return weatherData, nil
}
