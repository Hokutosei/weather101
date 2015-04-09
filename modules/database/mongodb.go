package database

import (
	"fmt"
	"log"
	"time"
	"weather101/modules/utilities"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var (
	session *mgo.Session

	dbName            = "weather_report101"
	weatherCollection = "weather"

	// set hours per query day
	week             time.Duration = 168
	twoDays          time.Duration = 48
	hoursPerDayQuery               = week
)

// StartMongoDb start mongodb instance
func StartMongoDb() {
	//currentSession, err := mgo.Dial("107.167.180.219:27017")
	currentSession, err := mgo.Dial("104.155.227.195:27020")
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

// SaveAndPrint save mongodb data and print output
func (w *WeatherData) SaveAndPrint(startTime time.Time, toPrint ...string) (bool, error) {
	sc := SessionCopy()
	c := sc.DB(dbName).C(weatherCollection)
	defer sc.Close()

	// w.CreatedAt = fmt.Sprintf("%v", time.Now().Local())
	w.CreatedAt = time.Now()
	err := c.Insert(w)
	if err != nil {
		fmt.Println(err)
		return false, err
	}

	endTime := fmt.Sprintf("took: %v", time.Since(startTime))
	toPrint = append(toPrint, endTime)
	utilities.InlinePrint(toPrint...)

	return true, nil
}

// getAllCities list down city names
func getAllCities() ([]Cities, error) {
	sc := SessionCopy()
	c := sc.DB(dbName).C(weatherCollection)
	defer sc.Close()

	result := []Cities{}
	query := []bson.M{{"$group": bson.M{
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
	c := sc.DB(dbName).C(weatherCollection)
	defer sc.Close()

	results := []WeatherData{}
	query := bson.M{"name": name}
	err := c.Find(query).All(&results)
	_ = err
	fmt.Println(len(results))
	fmt.Println(time.Since(start))
}

// GetWeatherData make http request getting data
func (w *WeatherData) GetWeatherData() {
	cities, _ := getAllCities()
	for _, city := range cities {
		go func(name string) {
			getAllWeatherByCity(name)
		}(city.Name)
	}
}

// GetIndex main index data getter
func (w *WeatherData) GetIndex(chanWeather chan []AggregateWeather) ([]AggregateWeather, error) {
	sc := SessionCopy()
	c := sc.DB(dbName).C(weatherCollection)
	defer sc.Close()

	// go GetWeatherData()

	//db.weather.aggregate([{$group: {_id: "$name", items: {$push: {temp: "$main.temp"}}}}])
	result := []AggregateWeather{}

	// aggregation query
	// group by name, sum, and
	// make an array of data that group by name
	gte := time.Now().Add(-time.Hour * hoursPerDayQuery)
	lte := time.Now()
	fmt.Println("query for this times gte: ", gte, " lte: ", lte)

	query := []bson.M{
		{"$match": bson.M{
			"created_at": bson.M{"$gte": gte, "$lte": lte},
		}},
		{"$group": bson.M{
			"_id": "$name",
			"sum": bson.M{"$sum": 1},
			"items": bson.M{
				"$push": bson.M{
					"temp":        "$main.temp",
					"created_at":  "$created_at",
					"description": "$weather"}},
		}},
	}

	start := time.Now()
	err := c.Pipe(query).All(&result)

	// make temp conversion here!
	// fmt.Println(result.ConvertKelvinToCent())
	convertedResult := TemperatureDataConvertion(result)

	// benchmark how much time it took
	fmt.Println("aggregate took: ", time.Since(start))
	if err != nil {
		fmt.Println(err)
		return result, err
	}
	chanWeather <- convertedResult
	return convertedResult, nil
}
