package http_controllers

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"weather101/modules/database"
	"weather101/modules/utilities"
)

// WeatherResponse struct for HTTP response
type WeatherResponse struct {
	Status int
	Data   []database.AggregateWeather
}

// Index http controller
func Index(w http.ResponseWriter, r *http.Request) {
	log.Println("index rendered...")
	indexTemplate := "index.html"
	t := template.New(indexTemplate).Delims("{{%", "%}}")
	// indexVars := IndexVars{}

	parsedTemplateStr := fmt.Sprintf("public/%s", indexTemplate)
	t, _ = t.ParseFiles(parsedTemplateStr)
	t.Execute(w, nil)
}

// GetIndex http request handler for index data
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
