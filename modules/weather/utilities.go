package weather

import (
	"fmt"
	"net/http"
	"time"
)

// httpGet main func for http get
func httpGet(city_url string) (*http.Response, error) {
	timeout := time.Duration(10 * time.Second)
	client := http.Client{
		Timeout: timeout,
	}

	response, err := client.Get(city_url)
	if err != nil {
		// do better error handlin	g here
		//panic(err)
		fmt.Println("city: ", city_url, ": err: ", err)
		var x *http.Response
		return x, err
	}

	return response, nil
}
