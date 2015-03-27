package http_controllers

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"weather101/modules/database"
	"weather101/modules/utilities"
)

type WeatherResponse struct {
	Status int
	Data   []database.AggregateWeather
}

func Index(w http.ResponseWriter, r *http.Request) {
	log.Println("index rendered...")
	indexTemplate := "index.html"
	t := template.New(indexTemplate).Delims("{{%", "%}}")
	// indexVars := IndexVars{}

	parsed_template_str := fmt.Sprintf("public/%s", indexTemplate)
	t, _ = t.ParseFiles(parsed_template_str)
	t.Execute(w, nil)
}

func GetIndex(w http.ResponseWriter, r *http.Request) {
	log.Println("GetIndex handled!")

	var weatherData database.WeatherData

//	weatherData.GetWeatherData()
	weathers, err := weatherData.GetIndex()
	if err != nil {
		fmt.Println(err)
		return
	}

	response := &WeatherResponse{Status: 200, Data: weathers}
	utilities.RespondObjectToJson(w, response)
}
