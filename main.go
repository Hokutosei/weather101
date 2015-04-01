package main

import (
	"fmt"
	"net/http"

	"weather101/modules/database"
	"weather101/modules/weather"
)

var (
	serverPort = ":8003"
)

func main() {
	initializeAssets()
	startRoutes()

	go database.StartMongoDb()
	go database.StartAerospikeDb()

	go weather.StartGettingWeather()

	fmt.Println("server is listening to -->>", serverPort)
	err := http.ListenAndServe(serverPort, nil)

	returnErrorHandler(err)
}
