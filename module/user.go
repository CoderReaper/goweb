package module

import (
	myredis "github.com/CoderReaper/goweb/third/redis"
)

type UserDBModle struct {
}

func Login(email string, pass string) bool {
	var pool = myredis.GetRedisPoolInstance()
	ret, err := pool.Get().Do("GET", email)
	if err != nil {
		
	}
	println(ret)
	return true
}
