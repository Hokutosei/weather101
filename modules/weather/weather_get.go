package weather

import (
	"fmt"
	"net/http"
)

var (
	api_url = "http://api.openweathermap.org/data/2.5/weather?q="
)

func getWeather(city ...string) {
	for _, name := range city {
		go func(name string) {
			fmt.Println(name)
		}(name)
	}
}

func StartGettingWeather() {
	cities := []string{ "akiruno-shi", "paranaque", "omiya-shi", "machida-shi" }
	getWeather(cities...)
}