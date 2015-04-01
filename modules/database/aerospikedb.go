package database

import (
	"fmt"
	aerospike "github.com/aerospike/aerospike-client-go"
)

var (
	ac *aerospike.Client
)

// StartAerospikeDb connect to Aerospikedb
func StartAerospikeDb() {
	fmt.Println("connection to Aerospikedb")
	client, err := aerospike.NewClient("104.155.230.215", 3000)
	if err != nil {
		fmt.Println("err connecting to aerospike db!: ", err)
		return
	}

	fmt.Println("connected to aerospikedb")
	ac = client
}

// GetAllCities fetch all city names
func GetAllCities() {
	fmt.Println("called!")
	key, err := aerospike.NewKey("weather101", "cities", "")
	if err != nil {
		fmt.Println(err)
		return
	}

	rec, err := ac.Get(nil, key)
	_ = err

	fmt.Println(rec)
}
