package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	_ "time"
	"weather101/modules/config"
	"weather101/modules/database"
	"weather101/modules/utilities"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var (
	session *mgo.Session

	dbName            = "weather_report101"
	weatherCollection = "weather"

	// set hours per query day

	mongodbClusterKey string = "mongodb_cluster1"
	selectLimit       int    = 3000
)

// GetMongodbCluster retrieve mongodb cluster host
func GetMongodbCluster(host chan string) {
	mongodbCluster, err := config.EtcdRawGetValue(mongodbClusterKey)
	if err != nil {
		panic(err)
	}

	host <- mongodbCluster
}

// StartMongoDb start mongodb instance
func StartMongoDb() {
	//currentSession, err := mgo.Dial("107.167.180.219:27017")
	mongodbCluster := make(chan string)
	go GetMongodbCluster(mongodbCluster)

	host := <-mongodbCluster
	currentSession, err := mgo.Dial(host)
	if err != nil {
		log.Println("err connecting to mongodb!")
		log.Println("error: ", err)
		return
	}
	fmt.Println("connected to mongos!")
	session = currentSession
}

// SessionCopy make a copy of current session
func SessionCopy() *mgo.Session {
	return session.Copy()
}

// startMigrate get all weather data records
func startMigrate() {
	sc := SessionCopy()
	c := sc.DB(dbName).C(weatherCollection)
	defer sc.Close()

	fmt.Println("will query...")
	var weather []database.WeatherData
	err := c.Find(bson.M{}).Sort("-_id").Limit(selectLimit).All(&weather)
	_ = err
	fmt.Println("got!: ", len(weather))
	migrateLoop(weather...)
}

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

// migrateLoop loop through all items
func migrateLoop(weather ...database.WeatherData) {
	var dataPoints DataPoints
	fmt.Println(len(weather))
	fmt.Println(len(dataPoints.DataPoint))
	for _, item := range weather {
		var points [][]interface{}
		fmt.Println(item.CreatedAt)

		points = append(points, []interface{}{item.CreatedAt.UnixNano() / 1000000, utilities.ConvertCelsius(item.Main.Temp), item.Name})

		item := DataPoint{
			Columns:   []string{"time", "temperature", "city"},
			Name:      "weather101",
			Fields:    points,
			Timestamp: item.CreatedAt.UnixNano() / 1000000,
			Precision: "s",
		}
		dataPoints.DataPoint = append(dataPoints.DataPoint, item)
	}
	BulkInsertToInfluxDb(dataPoints)
}

// BulkInsertToInfluxDb builk insert to influxdb
func BulkInsertToInfluxDb(weather DataPoints) {
	fmt.Println("will save!")
	url := "http://107.167.180.219:8086/db/weather101/series"
	fmt.Println("URL:>", url)

	mJSON, _ := json.Marshal(weather.DataPoint)
	contentReader := bytes.NewReader(mJSON)
	req, _ := http.NewRequest("POST", url, contentReader)

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", "InfluxDBClient")
	req.SetBasicAuth("jeanepaul", "jinpol")

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))
}

func main() {
	StartMongoDb()
	startMigrate()
}
