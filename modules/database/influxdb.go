package database

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	_ "io/ioutil"
	"net/http"
	_ "time"

	"weather101/modules/config"
	"weather101/modules/utilities"
)

var (
	influxdbAddress = "influxdbAddr"
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

// influxDbConfig influxdb config getter
func influxDbConfig(key chan string) (string, error) {
	value, err := config.EtcdRawGetValue(influxdbAddress)
	if err != nil {
		return "", err
	}

	return value, err
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

// influxDbURLConstruct return influxdb url
func influxDbURLConstruct() string {
	getURL := make(chan string)
	go influxDbConfig(getURL)

	url := <-getURL
	strURL := fmt.Sprintf("http://%s:8086/db/weather101/series", url)
	fmt.Println(strURL)

	return strURL
}

// BulkInsertToInfluxDb builk insert to influxdb
func BulkInsertToInfluxDb(weather DataPoints) (string, error) {
	url := influxDbURLConstruct()

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
	return msg, nil
}
