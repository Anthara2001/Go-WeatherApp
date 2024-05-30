package main

import (
	"net/http"
)

var openWeatherMapApi = "64d6472a8c3f759f80c62604cec8da06"

type WeatherData struct {
	City string `json:"name"`
	Main struct {
		Temp string `json:"temp_max"`
	} `json:"main"`
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/hello", hello)
	mux.HandleFunc("/weather/", func())

	s := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}
	s.ListenAndServe()
}
