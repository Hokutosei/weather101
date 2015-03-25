package weather

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	_ "net/http"
	"time"

	"weather101/modules/database"
	"weather101/modules/utilities"
)

var (
	api_url     = "http://api.openweathermap.org/data/2.5/weather?q="
	delay       = 59
	loopCounter = 0
)

func getWeather(city ...string) {
	for _, name := range city {
		go func(name string) {
			start := time.Now()
			city_url := fmt.Sprintf("%v%v", api_url, name)

			// http request
			response, err := httpGet(city_url)
			if err != nil {
				fmt.Println("has error: ", name)
				return
			}
			defer response.Body.Close()

			// read response
			contents, err := ioutil.ReadAll(response.Body)
			if err != nil {
				fmt.Println(err)
				return
			}

			// put response data in struct
			// unmarshal json response to struct
			var dat database.WeatherData
			if err := json.Unmarshal(contents, &dat); err != nil {
				fmt.Println(err)
				return
			}

			// fmt.Println(dat) # debug code
			temp_str := fmt.Sprintf("temp: %v C ", utilities.ConvertCelsius(dat.Main.Temp))
			toPrint := []string{
				temp_str,
				name,
			}

			// save weather data to mongodb!
			saved, err := dat.SaveAndPrint(start, toPrint...)
			if err != nil {
				fmt.Println(err)
				return
			}

			// dead code
			_ = saved

		}(name)
	}
}

func mainWeatherGetter() {
	cities := []string{"akiruno-shi", "paranaque", "omiya-shi", "machida-shi"}
	getWeather(cities...)
}

func StartGettingWeather() {
	// get some initial data from start
	// mainWeatherGetter()

	for i := range time.Tick(time.Second * time.Duration(delay)) {
		_ = i
		loopCounter++
		fmt.Println(time.Now().Format(time.RFC850), " counter: ", loopCounter)
		mainWeatherGetter()
	}
}
