package database

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	_ "io/ioutil"
	"net/http"
	_ "time"

	"weather101/modules/utilities"
)

// DataPoints structure for post data
type DataPoints struct {
	DataPoint []DataPoint
}

// DataPoint struct to use in request
type DataPoint struct {
	Columns   []string        `json:"columns"`
	Name      string          `json:"name"`
	Fields    [][]interface{} `json:"points"`
	Timestamp interface{}     `json:"timestamp"`
	Precision string          `json:"precision"`
}

// SaveToInfluxDB loop through all items
func (weather *WeatherData) SaveToInfluxDB() (string, error) {
	var dataPoints DataPoints

	if weather.Name == "" {
		err := errors.New("has nil name")
		return "error saving", err
	}
	var points [][]interface{}
	createdAt := weather.CreatedAt.UnixNano() / 1000000

	points = append(points, []interface{}{createdAt, utilities.ConvertCelsius(weather.Main.Temp), weather.Name})

	pointItem := DataPoint{
		Columns:   []string{"time", "temperature", "city"},
		Name:      weather.Name,
		Fields:    points,
		Timestamp: createdAt,
		Precision: "s",
	}
	dataPoints.DataPoint = append(dataPoints.DataPoint, pointItem)

	msg, err := BulkInsertToInfluxDb(dataPoints)
	return msg, err
}

// BulkInsertToInfluxDb builk insert to influxdb
func BulkInsertToInfluxDb(weather DataPoints) (string, error) {
	url := "http://107.167.180.219:8086/db/weather101/series"

	mJSON, _ := json.Marshal(weather.DataPoint)
	contentReader := bytes.NewReader(mJSON)
	req, _ := http.NewRequest("POST", url, contentReader)

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", "InfluxDBClient")
	req.SetBasicAuth("jeanepaul", "jinpol")

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		// panic(err)
		fmt.Println("did not save!")
		return "", err
	}
	defer resp.Body.Close()

	msg := fmt.Sprintf("influxDB response Status: %s", resp.Status)
	// fmt.Println("response Headers:", resp.Header)
	// body, _ := ioutil.ReadAll(resp.Body)
	// fmt.Println("response Body:", string(body))
	return msg, nil
}
