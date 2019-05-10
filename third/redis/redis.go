package myredis

import (
	"sync"
	"time"

	"github.com/gomodule/redigo/redis"
)

var redpool *redis.Pool

var one sync.Once

func GetRedisPoolInstance() *redis.Pool {
	one.Do(func() {
		redpool = redis.NewPool(Connect, 10)
	})
	return redpool
}

//Connect conn
func Connect() (redis.Conn, error) {
	con, err := redis.Dial("tcp", "35.236.107.65:6379",
		redis.DialPassword("sunmoon"),
		redis.DialDatabase(0),
		redis.DialConnectTimeout(60*time.Second),
		redis.DialReadTimeout(60*time.Second),
		redis.DialWriteTimeout(60*time.Second))
	if err != nil {
		return nil, err
	}
	return con, nil
}

func init() {
	redpool = GetRedisPoolInstance()
}
