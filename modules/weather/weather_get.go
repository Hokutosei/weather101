package weather

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	_ "net/http"
	"sync"
	"time"

	"weather101/modules/currency"
	"weather101/modules/database"
	"weather101/modules/utilities"
)

var (
	apiURL      = "http://api.openweathermap.org/data/2.5/weather?q="
	delay       = 59
	loopCounter = 0
)

// getWeather main event loop
// that get all weather data
func getWeather(city ...string) {
	fmt.Println(city)
	var wg sync.WaitGroup
	for _, name := range city {
		wg.Add(1)
		go func(name string) {
			start := time.Now()
			cityURL := fmt.Sprintf("%v%v", apiURL, name)

			// http request
			response, err := httpGet(cityURL)
			if err != nil {
				fmt.Println("has error: ", name)
				return
			}
			defer wg.Done()
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
			tempStr := fmt.Sprintf("temp: %v C ", utilities.ConvertCelsius(dat.Main.Temp))
			var weatherDescription string
			if len(dat.Weather) >= 1 {
				weatherDescription = dat.Weather[0].Description
			} else {
				weatherDescription = "weather description NIL"
			}

			// output message
			toPrint := []string{
				tempStr,
				weatherDescription,
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
	wg.Wait()

	// extra get modules
	//  get peso currency
	currency.GetPeso()
}

// mainWeatherGetter call main event weather loop
func mainWeatherGetter() {
	// make channel for city list result from redis
	cityListChan := make(chan []string)

	// from goroutine, get all city list from redis
	go database.GetAllCityList(cityListChan)

	// cities := []string{"akiruno-shi", "paranaque", "omiya-shi", "machida-shi", "akishima-shi"}
	// cities := <-cityListChan
	getWeather(<-cityListChan...)
}

// StartGettingWeather initialize weather getter and setter
func StartGettingWeather() {
	// get some initial data from start
	// mainWeatherGetter()

	// periodical main loop for GET && Save weather
	for i := range time.Tick(time.Second * time.Duration(delay)) {
		_ = i
		loopCounter++
		fmt.Println(time.Now().Format(time.RFC850), " counter: ", loopCounter)
		mainWeatherGetter()
	}
}
