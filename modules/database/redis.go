package database

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"strings"
	"time"

	"weather101/modules/config"
)

var (
	pool *redis.Pool

	redisKey string = "redisHost"
)

// getRedisHost redis host getter
func getRedisHost(redisHost chan string) {
	redis, err := config.EtcdRawGetValue(redisKey)
	if err != nil {
		panic(err)
	}
	redisHost <- redis
}

// StartRedis make connection to redis
func StartRedis() {
	fmt.Println("connecting to redis..")
	pool = newPool()

}

func newPool() *redis.Pool {
	redisHost := make(chan string)
	go getRedisHost(redisHost)

	host := <-redisHost
	return &redis.Pool{
		MaxIdle:     3,
		IdleTimeout: 10 * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", host)
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
