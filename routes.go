package main

import (
	"fmt"
	"net/http"

	"weather101/modules/http_controllers"
)

func startRoutes() {
	fmt.Println("starting routes..")

	http.HandleFunc("/", http_controllers.Index)
}
