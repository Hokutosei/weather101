package weather

import (
	"fmt"
	"net/http"
	"time"
)

func httpGet(city_url string) (*http.Response, error) {
	timeout := time.Duration(5 * time.Second)
	client := http.Client{
		Timeout: timeout,
	}

	response, err := client.Get(city_url)
	if err != nil {
		// do better error handling here
		//panic(err)
		fmt.Println(err)
		var x *http.Response
		return x, err
	}
	return response, nil
}
