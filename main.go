package main

import (
	"fmt"
	"net/http"
)

var (
	serverPort = ":8000"
)

func main() {
	initializeAssets()
	startRoutes()

	fmt.Println("server is listening to -->>", serverPort)
	err := http.ListenAndServe(serverPort, nil)

	returnErrorHandler(err)
}
