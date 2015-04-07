package database

import "fmt"

// TemperatureDataConvertion convert weather items
func TemperatureDataConvertion(aggregatedWeather []AggregateWeather) {
	for _, item := range aggregatedWeather {
		go func(item AggregateWeather) {
			for i := range item.Items {

				// get item via pointer
				n := &item.Items[i]

				// then modify it
				n.Celsius = convertKelvin(n.Temp)
			}
		}(item)
	}
}

func convertKelvin(kelvin float64) int {
	return int(kelvin - 273.15)
}

// ConvertKelvinToCent converts current temperature
// kelvin to centigrade
func (wa *AggregateWeather) ConvertKelvinToCent() {
	fmt.Println(len(wa.Items))
}
