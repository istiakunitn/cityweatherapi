package router

import (
	"github.com/gorilla/mux"
	"github.com/istiakunitn/cityweatherapi/apis"
)

func NewRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", apis.HomeHandler)
	r.HandleFunc("/api/weather", apis.WeatherHandler).Methods("GET")

	return r
}
