package main

import (
	"fmt"
	"net/http"

	"weather101/modules/httpControllers"
)

func startRoutes() {
	fmt.Println("starting routes..")

	httpControllers.MakeConnections()

	http.HandleFunc("/", httpControllers.Index)
	http.HandleFunc("/get_index", httpControllers.GetIndex)

	http.HandleFunc("/get_admin", httpControllers.AdminIndex)
}
