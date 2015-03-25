package main

import (
	"fmt"
	"net/http"

	"weather102/modules/http_controllers"
)

func startRoutes() {
	fmt.Println("starting routes..")

	http.HandleFunc("/", http_controllers.Index)
}
