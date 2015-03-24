package weather

import (
	"fmt"
	"net/http"
)

func httpGet(city_url string) *http.Response {
	response, err := http.Get(city_url)
	if err != nil {
		// do better error handling here
		panic(err)
	}
	return response
}