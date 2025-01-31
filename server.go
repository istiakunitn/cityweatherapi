package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/istiakunitn/cityweatherapi/router"
)

func main() {
	router := router.NewRouter()
	http.Handle("/", router)

	fmt.Println("Server running on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", handlers.LoggingHandler(os.Stdout, http.DefaultServeMux)))
}
