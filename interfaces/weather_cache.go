package interfaces

import (
	"time"
)

type WeatherCacheEntry struct {
	Data      OpenWeatherResponse
	ExpiresAt time.Time
}
