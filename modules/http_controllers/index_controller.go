package http_controllers

import (
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
	chanWeather := make(chan []database.AggregateWeather)
	weathers, err := weatherData.GetIndex(chanWeather)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(len(weathers))

	for {
		// Blocks until a message is read
		_, msg, err := conn.ReadMessage()
		if err != nil {
			delete(connections, conn)
			conn.Close()
			return
		}
		log.Println(string(msg))

		for _, item := range <-chanWeather {
			sendAll([]byte{item})
		}
	}

	// var weatherData database.WeatherData
	//
	// //	weatherData.GetWeatherData()
	// weathers, err := weatherData.GetIndex()
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	//
	// response := &WeatherResponse{Status: 200, Data: weathers}
	// utilities.RespondObjectToJson(w, response)
}

func sendAll(msg []byte) {
	for conn := range connections {
		if err := conn.WriteMessage(websocket.TextMessage, msg); err != nil {
			delete(connections, conn)
			conn.Close()
		}
	}
}

func MakeConnections() {
	connections = make(map[*websocket.Conn]bool)
}

// func wsHandler(w http.ResponseWriter, r *http.Request) {
// 	// Taken from gorilla's website
// 	conn, err := websocket.Upgrade(w, r, nil, 1024, 1024)
// 	if _, ok := err.(websocket.HandshakeError); ok {
// 		http.Error(w, "Not a websocket handshake", 400)
// 		return
// 	} else if err != nil {
// 		log.Println(err)
// 		return
// 	}
// 	log.Println("Succesfully upgraded connection")
// 	connections[conn] = true
//
// 	for {
// 		// Blocks until a message is read
// 		_, msg, err := conn.ReadMessage()
// 		if err != nil {
// 			delete(connections, conn)
// 			conn.Close()
// 			return
// 		}
// 		log.Println(string(msg))
// 		sendAll(msg)
// 	}
// }
