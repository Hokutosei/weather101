package weather

import (
	_ "fmt"
	"net/http"
)

type WeatherData struct {
	Base   string `json:"base"`
	Clouds struct {
		All float64 `json:"all"`
	} `json:"clouds"`
	Cod   float64 `json:"cod"`
	Coord struct {
		Lat float64 `json:"lat"`
		Lon float64 `json:"lon"`
	} `json:"coord"`
	Dt   float64 `json:"dt"`
	ID   float64 `json:"id"`
	Main struct {
		Humidity float64 `json:"humidity"`
		Pressure float64 `json:"pressure"`
		Temp     float64 `json:"temp"`
		TempMax  float64 `json:"temp_max"`
		TempMin  float64 `json:"temp_min"`
	} `json:"main"`
	Name string `json:"name"`
	Sys  struct {
		Country string  `json:"country"`
		ID      float64 `json:"id"`
		Message float64 `json:"message"`
		Sunrise float64 `json:"sunrise"`
		Sunset  float64 `json:"sunset"`
		Type    float64 `json:"type"`
	} `json:"sys"`
	Weather []struct {
		Description string  `json:"description"`
		Icon        string  `json:"icon"`
		ID          float64 `json:"id"`
		Main        string  `json:"main"`
	} `json:"weather"`
	Wind struct {
		Deg   float64 `json:"deg"`
		Gust  float64 `json:"gust"`
		Speed float64 `json:"speed"`
	} `json:"wind"`
}



func httpGet(city_url string) *http.Response {
	response, err := http.Get(city_url)
	if err != nil {
		// do better error handling here
		panic(err)
	}
	return response
}