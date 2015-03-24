package weather

import (
	"fmt"
	_ "net/http"
	"time"
	"io/ioutil"
	"encoding/json"
)

var (
	api_url = "http://api.openweathermap.org/data/2.5/weather?q="
	delay = 4
)

func getWeather(city ...string) {
	for _, name := range city {
		go func(name string) {
			city_url := fmt.Sprintf("%v%v", api_url, name)

			// http request
			response := httpGet(city_url)
			defer response.Body.Close()

			// read response
			contents, err := ioutil.ReadAll(response.Body)
			if err!= nil {
				fmt.Println(err)
				return
			}

			// put response data in struct
			var dat WeatherData
			if err := json.Unmarshal(contents, &dat); err != nil {
				fmt.Println(err)
				return
			}

			fmt.Println(dat)

			fmt.Println(response)

		}(name)
	}
}

func StartGettingWeather() {

	for i := range time.Tick(time.Second * time.Duration(delay)) {
		_ = i
		cities := []string{ "akiruno-shi", "paranaque", "omiya-shi", "machida-shi" }
		getWeather(cities...)		
	}
}