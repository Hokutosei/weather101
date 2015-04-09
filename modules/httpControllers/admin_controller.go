package httpControllers

import (
	"fmt"
	"net/http"

	"weather101/modules/database"
	"weather101/modules/utilities"
)

// AdminIndexResponse data struct for json response
type AdminIndexResponse struct {
	Status int
	Data   interface{}
}

// CityResponse data struct for json response
type CityResponse struct {
	CityList []string `json:"city_list"`
}

// AdminIndex admin page handler
func AdminIndex(w http.ResponseWriter, r *http.Request) {
	fmt.Println("AdminIndex handled!")
	cityList := make(chan []string)

	go database.GetAllCityList(cityList)

	responseData := &CityResponse{CityList: <-cityList}
	response := &AdminIndexResponse{Status: 200, Data: responseData}
	utilities.RespondObjectToJson(w, response)
}
