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

	// w.CreatedAt = fmt.Sprintf("%v", time.Now().Local())
	w.CreatedAt = time.Now()
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
