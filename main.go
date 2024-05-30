package main

import (
	"encoding/json"
	"net/http"
	"strings"
)

var openWeatherMapApiKey = "64d6472a8c3f759f80c62604cec8da06"

type WeatherData struct {
	City string `json:"name"`
	Main struct {
		Temp string `json:"temp_max"`
	} `json:"main"`
}

func hello(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("Welcome to WeatherApp by Golang"))
}

func query(city string) WeatherData {
	var d WeatherData
	resp, _ := http.Get("https://api.openweathermap.org/data/2.5/weather?q=" + city + "&appid=" + openWeatherMapApiKey)
	json.NewDecoder(resp.Body).Decode(&d)
	return d
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/hello", hello)
	mux.HandleFunc("/weather/",

		func(res http.ResponseWriter, req *http.Request) {
			city := strings.SplitN(req.URL.Path, "/", 3)[2]
			data := query(city)
			json.NewEncoder(res).Encode(data)
		},
	)

	s := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}
	s.ListenAndServe()
}
