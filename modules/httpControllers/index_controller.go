package httpControllers

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"

	"weather101/modules/database"
	_ "weather101/modules/utilities"

	"github.com/gorilla/websocket"
)

var (
	connections map[*websocket.Conn]bool
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

	//websocket
	conn, err := websocket.Upgrade(w, r, nil, 1024, 1024)
	if _, ok := err.(websocket.HandshakeError); ok {
		http.Error(w, "not a websocket handshake", 400)
		return
	} else if err != nil {
		log.Println(err)
		return
	}

	log.Println("successfully upgrade connection!")
	connections[conn] = true

	var weatherData database.WeatherData

	//	weatherData.GetWeatherData()
	chanWeather := make(chan database.AggregateWeather)

	// query and analyze weather data
	go weatherData.GetIndex(chanWeather)

	// long poll query data func and send update
	go longPollWeather(chanWeather)

	for {
		encodedData, err := json.Marshal(<-chanWeather)
		_ = err
		sendAll(encodedData)
	}
}

func sendAll(msg []byte) {
	for conn := range connections {
		if err := conn.WriteMessage(websocket.TextMessage, msg); err != nil {
			delete(connections, conn)
			conn.Close()
		}
	}
}

// MakeConnections initialize connections
func MakeConnections() {
	connections = make(map[*websocket.Conn]bool)
}
