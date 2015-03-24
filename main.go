package main

import (
	"fmt"
	"net/http"

	"web102/modules/database"
)

var (
	serverPort = ":8000"
)

func main() {
	initializeAssets()
	startRoutes()

	go database.StartMongoDb()

	fmt.Println("server is listening to -->>", serverPort)
	err := http.ListenAndServe(serverPort, nil)

	returnErrorHandler(err)
}
