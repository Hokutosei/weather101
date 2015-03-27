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

func getAllCities() ([]Cities, error) {
	sc := SessionCopy()
	c := sc.DB(dbName).C(weather_collection)
	defer sc.Close()

	result := []Cities {}
	query := []bson.M{{ "$group": bson.M{
			"_id": "$name",
		}},
	}
	err := c.Pipe(query).All(&result)
	if err != nil {
		return result, err
	}
	return result, nil
}

func getAllWeatherByCity(name string) {
	start := time.Now()
	sc := SessionCopy()
	c := sc.DB(dbName).C(weather_collection)
	defer sc.Close()

	results := []WeatherData{}
	query := bson.M{"name": name}
	err := c.Find(query).All(&results)
	_ = err
	fmt.Println(len(results))
	fmt.Println(time.Since(start))
}

func (w *WeatherData) GetWeatherData() {
	cities, _ := getAllCities()
	for _, city := range cities {
		go func(name string) {
			getAllWeatherByCity(name)
		}(city.Name)
	}
}

func (w *WeatherData) GetIndex() ([]AggregateWeather, error) {
	sc := SessionCopy()
	c := sc.DB(dbName).C(weather_collection)
	defer sc.Close()

	// go GetWeatherData()

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
