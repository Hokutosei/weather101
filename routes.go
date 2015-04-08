package main

import (
	"fmt"
	"net/http"

	"weather101/modules/http_controllers"
)

func startRoutes() {
	fmt.Println("starting routes..")

	http_controllers.MakeConnections()

	http.HandleFunc("/", http_controllers.Index)
	http.HandleFunc("/get_index", http_controllers.GetIndex)

	http.HandleFunc("/get_admin", http_controllers.AdminIndex)
}
