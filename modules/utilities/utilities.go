package utilities

import (
	"encoding/json"
	"fmt"
	"net/http"
)

var (
	kelvin        float64 = 272.15
	paddingMaxLen         = 15
	paddingMinLen         = 5
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

	fmt.Println("---------------->> ", printStr)
}

// transform string for aesthetic and readable in CLI
func StringifyAndPadding(str interface{}) string {
	var paddedStr string

	initStr := fmt.Sprintf("%v ", str)
	padMaxLen := paddingMaxLen

	if len(initStr) < paddingMinLen {
		padMaxLen = paddingMinLen
	}

	// add padding
	for i := 0; i < (padMaxLen - len(initStr)); i++ {
		_ = i
		paddedStr += " "
	}

	paddedStr += initStr + " | "
	return paddedStr
}

func RespondObjectToJson(w http.ResponseWriter, object interface{}) {
	js, err := json.Marshal(object)
	if err != nil {
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

//func ReadRequestPostRequest(r *http.Request) (database.UserEditParams, error) {
//	query := r.URL.Query()
//	fmt.Println(query["nickname"][0])
//	decoder := json.NewDecoder(r.Body)
//
//	var data database.UserEditParams
//	err := decoder.Decode(&data)
//
//	if err != nil {
//		fmt.Println(err)
//		return data, err
//	}
//	return data, nil
//}
