package myredis

import (
	"sync"
	"time"

	"github.com/gomodule/redigo/redis"
)

var redpool *redis.Pool

var one sync.Once

func getRedisPoolInstance() *redis.Pool {
	one.Do(func() {
		redpool = redis.NewPool(connect, 10)
	})
	return redpool
}

//redis 命令
const (
	//key
	//DEL删除某个key
	DEL = "DEL"
	//EXIST 是否存在key
	EXIST = "EXIST"
	//EXPIRE 设置key的过期时间
	EXPIRE = "EXPIRE"
	//EXPIREAT key timestamp 设置key的过期时间时间戳
	EXPIREAT = "EXPIREAT"
	//PEXPIRE key milliseconds 设置key的过期时间时间戳
	PEXPIRE = "PEXPIRE"
	//PEXPIREAT key milliseconds-timestamp
	PEXPIREAT = "PEXPIREAT"
	//KEYS pattern 查找所有符合给定模式( pattern)的 key
	KEYS = "KEYS"
	//MOVE key db 将当前数据库的 key 移动到给定的数据库 db 当中
	MOVE = "MOVE"
	//PTTL key 以毫秒为单位返回 key 的剩余的过期时间
	PTTL = "PTTL"
	//TTL key 以秒为单位返回 key 的剩余的过期时间
	TTL = "TTL"
	//RANDOMKEY 从当前数据库中随机返回一个 key
	RANDOMKEY = "RANDOMKEY"
	//RENAME key newkey 修改 key 的名称
	RENAME = "RENAME"
	//TYPE key 返回 key 所储存的值的类型
	TYPE = "TYPE"

	// STRING 字符串
	//SET key value 设置指定 key 的值
	SET = "SET"
	//GET key 获取指定 key 的值。
	GET = "GET"
	//GETRANGE key start end 返回 key 中字符串值的子字符串
	GETRANGE = "GETRANGE"
	//GETSET key value将给定 key 的值设为 value ，并返回 key 的旧值(old value)
	GETSET = "GETSET"
	//MGET key1 [key2..] 获取所有(一个或多个)给定 key 的值
	MGET = "MGET"
	//SETNX key value只有在 key 不存在时设置 key 的值
	SETNX = "SETNX"
	//STRLEN key返回 key 所储存的字符串值的长度
	STRLEN = "STRLEN"
	//MSET key value [key value ...]同时设置一个或多个 key-value 对
	MSET = "MSET"
	//MSETNX key value [key value ...] 同时设置一个或多个 key-value 对，当且仅当所有给定 key 都不存在
	MSETNX = "MSETNX"
	//SETEX key seconds value将值 value 关联到 key ，并将 key 的过期时间设为 seconds (以秒为单位)
	SETEX = "SETEX"
	//PSETEX key milliseconds value这个命令和 SETEX 命令相似，但它以毫秒为单位设置 key 的生存时间，而不是像 SETEX 命令那样，以秒为单位
	PSETEX = "PSETEX"
	//INCR key将 key 中储存的数字值增一
	INCR = "INCR"
	//INCRBY key increment将 key 所储存的值加上给定的增量值（increment
	INCRBY = "INCRBY"
	//INCRBYFLOAT key increment将 key 所储存的值加上给定的浮点增量值（increment）
	INCRBYFLOAT = "INCRBYFLOAT"
	//DECR key将 key 中储存的数字值减一
	DECR = "DECR"
	//	DECRBY key decrement key 所储存的值减去给定的减量值（decrement
	DECRBY = "DECRBY"
	//APPEND key value如果 key 已经存在并且是一个字符串， APPEND 命令将指定的 value 追加到该 key 原来值（value）的末尾
	APPEND = "APPEND"
	// HASH 哈希表
)

func connect() (redis.Conn, error) {
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
	redpool = getRedisPoolInstance()
}

//GetConn 获取连接
func GetConn() redis.Conn {
	return redpool.Get()
}

//PutBack 放回池子
func PutBack(conn redis.Conn) {
	if conn != nil {
		conn.Close()
	}
}

//Do do
func Do(commandName string, args ...interface{}) (reply interface{}, err error) {
	conn := GetConn()
	reply, err = conn.Do(commandName, args...)
	defer PutBack(conn)
	return
}
