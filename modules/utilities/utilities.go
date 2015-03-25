package utilities

import (
	"fmt"
)

var (
	kelvin float64 = 272.15
	paddingMaxLen = 12
	paddingMinLen = 5
)

func ConvertCelsius(temp float64) int {
	result := temp - kelvin
	return int(result)
}

func InlinePrint(toPrint ...string) {
	var printStr string
	for _, str := range toPrint {
		printStr += StringifyAndPadding(str)
	}

	fmt.Println("------------------------------ ", printStr)
}

func StringifyAndPadding(str interface{}) string {
	initStr := fmt.Sprintf("%v ", str)
	padMaxLen := paddingMaxLen

	var paddedStr string
	
	if len(initStr) < paddingMinLen {
		padMaxLen = paddingMinLen
	}

	// add padding
	for i := 0; i < (padMaxLen - len(initStr)); i++ {
		_ = i
		paddedStr += " "
	}

	paddedStr += initStr
	return paddedStr
}