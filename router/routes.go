package router

import (
	"github.com/gorilla/mux"
	"github.com/istiakunitn/cityweatherapi/api"
	"github.com/istiakunitn/cityweatherapi/pages"
)

func NewRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", pages.HomeHandler)
	r.HandleFunc("/api/weather", api.WeatherHandler).Methods("GET")

	return r
}
