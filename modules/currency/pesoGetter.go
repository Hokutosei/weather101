package currency

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"time"
	"weather101/modules/database"
	"weather101/modules/utilities"
)

// GetPeso GET request for getting peso conversion
func GetPeso() {
	start := time.Now()
	getPesoURL := "http://www.freecurrencyconverterapi.com/api/v3/convert?q=JPY_PHP&compact=y"
	resp, err := utilities.HTTPGet(getPesoURL)

	if err != nil {
		fmt.Println("err getpeso, ", err)
		return
	}
	defer resp.Body.Close()

	contents, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("err getpeso ioutil: ", err)
		return
	}

	var pesoCurrency database.PesoToYen
	if err := json.Unmarshal(contents, &pesoCurrency); err != nil {
		fmt.Println("err json unmarshal: ", err)
		return
	}

	toPrint := []string{
		fmt.Sprintf("yen -> peso: %v", pesoCurrency.JPY_PHP.Val),
	}

	saved, err := pesoCurrency.PesoSaveAndPrint(start, toPrint...)
	if err != nil {
		fmt.Println("pesocurrency not saved: ", err)
	}

	_ = saved
}
