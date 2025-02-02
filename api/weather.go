package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/istiakunitn/cityweatherapi/cache"
	"github.com/istiakunitn/cityweatherapi/interfaces"
	"github.com/istiakunitn/cityweatherapi/services"
)

func WeatherHandler(w http.ResponseWriter, r *http.Request) {
	city := r.URL.Query().Get("city")
	if city == "" {
		http.Error(w, "City parameter is required", http.StatusBadRequest)
		return
	}

	weatherCache := cache.GetCacheInstance[interfaces.OpenWeatherResponse]()

	data, found := weatherCache.Get(city)

	if found {
		fmt.Println("Cached Weather Data:", data)
	} else {
		fmt.Println("Cache miss!")
	}

	weatherData, err := services.GetFromOpenWeather(city)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	weatherCache.Set(city, weatherData, 5*time.Minute)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(weatherData)
}
