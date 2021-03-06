package database

import "fmt"

// TemperatureDataConvertion convert weather items
func TemperatureDataConvertion(aggregatedWeather []AggregateWeather, modifiedWeatherData chan AggregateWeather) {
	for _, item := range aggregatedWeather {
		go func(item AggregateWeather) {
			for i := range item.Items {

				// get item via pointer
				n := &item.Items[i]

				// then modify it
				resultChan := make(chan int)
				//n.Celsius = convertKelvin(n.Temp, resultChan)
				go convertKelvin(n.Temp, resultChan)
				out := <-resultChan
				n.Celsius = out
			}
		}(item)
		modifiedWeatherData <- item
	}
}

func convertKelvin(kelvin float64, resultInt chan int) {
	// return result
	resultInt <- int(kelvin - 273.15)
}

// ConvertKelvinToCent converts current temperature
// kelvin to centigrade
func (wa *AggregateWeather) ConvertKelvinToCent() {
	fmt.Println(len(wa.Items))
}
