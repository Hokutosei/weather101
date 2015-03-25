package utilities

import (
	"fmt"
)

var (
	kelvin float64 = 272.15
)

func ConvertCelsius(temp float64) int {
	result := temp - kelvin
	return int(result)
}

func InlinePrint(toPrint ...string) {
	var printStr string
	for _, str := range toPrint {
		printStr += Stringify(str)
	}

	// fmt.Println("saved!")
	// fmt.Println("------------------------------", name, " took: ", time.Since(start_time))
	fmt.Println("------------------------------ ", printStr)
}

func Stringify(str interface{}) string {
	initStr := fmt.Sprintf("%v ", str)
	padMaxLen := 15
	
	var paddedStr string
	
	if len(initStr) < 5 {
		padMaxLen = 5
	}


	for i := 0; i < (padMaxLen - len(initStr)); i++ {
		_ = i
		paddedStr += " "
	}
	paddedStr += initStr

	return paddedStr
}