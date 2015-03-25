package database

import (
	"fmt"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
	"time"

	"weather101/modules/utilities"
)

var (
	session *mgo.Session

	dbName             = "weather_report101"
	weather_collection = "weather"
)

type WeatherData struct {
	CreatedAt string `bson:"created_at" json:"created_at"`
	Base      string `json:"base"`
	Clouds    struct {
		All float64 `json:"all"`
	} `json:"clouds"`
	Cod   float64 `json:"cod"`
	Coord struct {
		Lat float64 `json:"lat"`
		Lon float64 `json:"lon"`
	} `json:"coord"`
	Dt   float64 `json:"dt"`
	ID   float64 `json:"id"`
	Main struct {
		Humidity float64 `json:"humidity"`
		Pressure float64 `json:"pressure"`
		Temp     float64 `json:"temp"`
		TempMax  float64 `json:"temp_max"`
		TempMin  float64 `json:"temp_min"`
	} `json:"main"`
	Name string `json:"name"`
	Sys  struct {
		Country string  `json:"country"`
		ID      float64 `json:"id"`
		Message float64 `json:"message"`
		Sunrise float64 `json:"sunrise"`
		Sunset  float64 `json:"sunset"`
		Type    float64 `json:"type"`
	} `json:"sys"`
	Weather []struct {
		Description string  `json:"description"`
		Icon        string  `json:"icon"`
		ID          float64 `json:"id"`
		Main        string  `json:"main"`
	} `json:"weather"`
	Wind struct {
		Deg   float64 `json:"deg"`
		Gust  float64 `json:"gust"`
		Speed float64 `json:"speed"`
	} `json:"wind"`
}

func StartMongoDb() {
	current_session, err := mgo.Dial("107.167.180.219:27017")
	if err != nil {
		log.Println("err connecting to mongodb!")
		log.Println("error: ", err)
		return
	}
	fmt.Println("connected to mongodb!")
	session = current_session
}

func SessionCopy() *mgo.Session {
	return session.Copy()
}

func (w *WeatherData) SaveAndPrint(start_time time.Time, toPrint ...string) (bool, error) {
	sc := SessionCopy()
	c := sc.DB(dbName).C(weather_collection)
	defer sc.Close()

	w.CreatedAt = fmt.Sprintf("%v", time.Now().Local())
	err := c.Insert(w)
	if err != nil {
		fmt.Println(err)
		return false, err
	}

	endTime := fmt.Sprintf("took: %v", time.Since(start_time))
	toPrint = append(toPrint, endTime)
	utilities.InlinePrint(toPrint...)

	return true, nil
}

type AggregateWeather struct {
	Name  string `bson:"_id"`
	Sum   int
	Items []struct {
		Temp      float64 `json:"temp"`
		CreatedAt string  `bson:"created_at" json:"created_at"`
	}
}

func (w *WeatherData) GetIndex() ([]AggregateWeather, error) {
	sc := SessionCopy()
	c := sc.DB(dbName).C(weather_collection)
	defer sc.Close()

	//db.weather.aggregate([{$group: {_id: "$name", items: {$push: {temp: "$main.temp"}}}}])
	result := []AggregateWeather{}

	// aggregation query
	// group by name, sum, and
	// make an array of data that group by name
	query := []bson.M{{"$group": bson.M{
		"_id": "$name",
		"sum": bson.M{"$sum": 1},
		"items": bson.M{
			"$push": bson.M{
				"temp":       "$main.temp",
				"created_at": "$created_at"}},
	}},
	}

	start := time.Now()
	err := c.Pipe(query).All(&result)

	// benchmark how much time it took
	fmt.Println("aggregate took: ", time.Since(start))
	if err != nil {
		fmt.Println(err)
		return result, err
	}
	return result, nil
}
