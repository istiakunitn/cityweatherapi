package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	OpenWeatherBaseUrl string `mapstructure:"OPEN_WEATHER_BASE_URL"`
	OpenWeatherApiKey  string `mapstructure:"OPEN_WEATHER_API_KEY"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
