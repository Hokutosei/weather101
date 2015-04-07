package database

import "fmt"

// TemperatureDataConvertion convert weather items
func TemperatureDataConvertion(aggregatedWeather []AggregateWeather) []AggregateWeather {
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
	}
	return aggregatedWeather
}

func convertKelvin(kelvin float64, resultInt chan int) {
	result := int(kelvin - 273.15)
	if result == 0 {
		panic("error")
	}
	// return result
	resultInt <- result
}

// ConvertKelvinToCent converts current temperature
// kelvin to centigrade
func (wa *AggregateWeather) ConvertKelvinToCent() {
	fmt.Println(len(wa.Items))
}
