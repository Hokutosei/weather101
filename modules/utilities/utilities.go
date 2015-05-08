package utilities

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

var (
	kelvin        float64 = 272.15
	paddingMaxLen         = 15
	paddingMinLen         = 5
)

// ConvertCelsius convert temp to celsius
func ConvertCelsius(temp float64) int {
	result := temp - kelvin
	return int(result)
}

// InlinePrint print messages inline,
// toPrint message slice
func InlinePrint(toPrint ...string) {
	var printStr string
	for _, str := range toPrint {
		printStr += StringifyAndPadding(str)
	}

	fmt.Println("---------------->> ", printStr)
}

// StringifyAndPadding transform string for aesthetic and readable in CLI
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

// RespondObjectToJson send response objects/structs to json
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

// HTTPGet an http GET request with timeout
func HTTPGet(url string) (*http.Response, error) {
	timeout := time.Duration(5 * time.Second)
	client := http.Client{
		Timeout: timeout,
	}

	response, err := client.Get(url)
	if err != nil {
		// do better error handling here
		//panic(err)
		fmt.Println(err)
		var x *http.Response
		return x, err
	}
	return response, nil
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
