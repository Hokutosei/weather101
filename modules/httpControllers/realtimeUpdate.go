package httpControllers

import (
	"fmt"
	"time"
	"weather101/modules/database"
)

var (
	secondsCounter time.Duration = 120
)

func longPollWeather(longPollChan chan database.AggregateWeather) {
	for i := range time.Tick(time.Second * secondsCounter) {
		fmt.Println(i)
		var weatherData database.WeatherData

		//	weatherData.GetWeatherData()
		chanWeather := make(chan database.AggregateWeather)

		// query and analyze weather data
		go weatherData.GetIndex(chanWeather)

		go func() {
			for {

				out := <-chanWeather

				longPollChan <- out

			}
		}()
	}
}
