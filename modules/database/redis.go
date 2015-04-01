package database

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"strings"
	"time"
)

var (
	pool *redis.Pool
)

// StartRedis make connection to redis
func StartRedis() {
	fmt.Println("connecting to redis..")
	pool = newPool()

}

func newPool() *redis.Pool {
	return &redis.Pool{
		MaxIdle:     3,
		IdleTimeout: 10 * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", "130.211.255.234:6379")
			if err != nil {
				return nil, err
			}
			// if _, err := c.Do("AUTH", nil); err != nil {
			//     c.Close()
			//     return nil, err
			// }
			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}
}

// keyConstructor returns a concatenated or joined slice with
// colon string separator
func keyConstructor(stringName ...string) string {
	return strings.Join(stringName, ":")
}

// GetAllCityList query all city names
func GetAllCityList(cityList chan []string) {
	r := pool.Get()
	defer r.Close()

	strName := []string{"city", "list"}
	key := keyConstructor(strName...)

	resp, err := redis.Strings(r.Do("LRANGE", key, "0", "-1"))
	if err != nil {
		fmt.Println(err)
		cityList <- []string{}
	}

	cityList <- resp
}
