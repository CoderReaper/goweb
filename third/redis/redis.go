package myredis

import (
	"strconv"
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
	//HDEL key field1 [field2] 删除一个或多个哈希表字段
	HDEL = "HDEL"
	//HGET key field 获取存储在哈希表中指定字段的值
	HGET = "HGET"
	//HGETALL key 获取在哈希表中指定 key 的所有字段和值
	HGETALL = "HGETALL"
	//HINCRBY key field increment 为哈希表 key 中的指定字段的整数值加上增量 increment
	HINCRBY = "HINCRBY"
	//HINCRBYFLOAT key field increment 为哈希表 key 中的指定字段的浮点数值加上增量 increment
	HINCRBYFLOAT = "HINCRBYFLOAT"
	//HKEYS key 获取所有哈希表中的字段
	HKEYS = "HKEYS"
	//HLEN key 获取哈希表中字段的数量
	HLEN = "HLEN"
	//HMGET key field1 [field2] 获取所有给定字段的值
	HMGET = "HMGET"
	//HMSET key field1 value1 [field2 value2 ] 同时将多个 field-value (域-值)对设置到哈希表 key 中
	HMSET = "HMSET"
	//HSET key field value 将哈希表 key 中的字段 field 的值设为 value
	HSET = "HSET"
	//HSETNX key field value 只有在字段 field 不存在时，设置哈希表字段的值
	HSETNX = "HSETNX"
	//HVALS key 获取哈希表中所有值
	HVALS = "HVALS"
	//HSCAN key cursor [MATCH pattern] [COUNT count] 迭代哈希表中的键值对
	HSCAN = "HSCAN"
	//LIST
	//BLPOP key1 [key2 ] timeout 移出并获取列表的第一个元素， 如果列表没有元素会阻塞列表直到等待超时或发现可弹出元素为止
	BLPOP = "BLPOP"
	//	BRPOP key1 [key2 ] timeout 移出并获取列表的最后一个元素， 如果列表没有元素会阻塞列表直到等待超时或发现可弹出元素为止
	BRPOP = "BRPOP"
	//LINDEX key index 通过索引获取列表中的元素
	LINDEX = "LINDEX"
	//LINSERT key BEFORE|AFTER pivot value 在列表的元素前或者后插入元素
	LINSERT = "LINSERT"
	//LLEN key 获取列表长度
	LLEN = "LLEN"
	//LPOP key 移出并获取列表的第一个元素
	LPOP = "LPOP"
	//LPUSH key value1 [value2] 将一个或多个值插入到列表头部
	LPUSH = "LPUSH"
	//LPUSHX key value 将一个值插入到已存在的列表头部
	LPUSHX = "LPUSHX"
	//LRANGE key start stop 获取列表指定范围内的元素
	LRANGE = "LRANGE"
	//LREM key count value 移除列表元素
	LREM = "LREM"
	//LSET key index value 通过索引设置列表元素的值
	LSET = "LSET"
	//RPOP key 移除列表的最后一个元素，返回值为移除的元素
	RPOP = "RPOP"
	//RPOPLPUSH source destination 移除列表的最后一个元素，并将该元素添加到另一个列表并返回
	RPOPLPUSH = "RPOPLPUSH"
	//RPUSH key value1 [value2] 在列表中添加一个或多个值
	RPUSH = "RPUSH"
	//RPUSHX key value 为已存在的列表添加值
	RPUSHX = "RPUSHX"

	//SET
	//	SADD key member1 [member2] 向集合添加一个或多个成员
	SADD = "SADD"
	//	SCARD key 获取集合的成员数
	SCARD = "SCARD"
	//SDIFF key1 [key2] 返回给定所有集合的差集
	SDIFF = "SDIFF"
	//	SDIFFSTORE destination key1 [key2] 返回给定所有集合的差集并存储在 destination 中
	SDKFFSTORE = "SDIFFSTORE"
	//SINTER key1 [key2] 返回给定所有集合的交集
	SINTER = "SINTER"
	//SINTERSTORE destination key1 [key2] 返回给定所有集合的交集并存储在 destination 中
	SINTERSTORE = "SINTERSTORE"
	//SISMEMBER key member 判断 member 元素是否是集合 key 的成员
	SISMEMBER = "SISMEMBER"
	//SMEMBERS key 返回集合中的所有成员
	SMEMBERS = "SMEMBERS"
	//SMOVE source destination member 将 member 元素从 source 集合移动到 destination 集合
	SMOVE = "SMOVE"
	//	SPOP key 移除并返回集合中的一个随机元素
	SPOP = "SPOP"
	//	SRANDMEMBER key [count] 返回集合中一个或多个随机数
	SRANDMEMBER = "SRANDMEMBER"
	//	SREM key member1 [member2] 移除集合中一个或多个成员
	SREM = "SREM"
	//SUNION key1 [key2] 返回所有给定集合的并集
	SUNION = "SUNION"
	//	SUNIONSTORE destination key1 [key2] 所有给定集合的并集存储在 destination 集合中
	SUNIONSTORE = "SUNIONSTORE"
	//SSCAN key cursor [MATCH pattern] [COUNT count] 迭代集合中的元素
	SSCAN = "SSCAN"

	//ZADD key score1 member1 [score2 member2] 向有序集合添加一个或多个成员，或者更新已存在成员的分数
	ZADD = "ZADD"
	//ZCARD key 获取有序集合的成员数
	ZCARD = "ZCARD"
	//	ZCOUNT key min max 计算在有序集合中指定区间分数的成员数
	ZCOUNT = "ZCOUNT"
	//	ZINCRBY key increment member 有序集合中对指定成员的分数加上增量 increment
	ZINCRBY = "ZINCRBY"
	//	ZINTERSTORE destination numkeys key [key ...] 计算给定的一个或多个有序集的交集并将结果集存储在新的有序集合 key 中
	ZINTERSTORE = "ZINTERSTORE"
	//ZLEXCOUNT key min max 在有序集合中计算指定字典区间内成员数量
	ZLEXCOUNT = "ZLEXCOUNT"
	//ZRANGE key start stop [WITHSCORES] 通过索引区间返回有序集合成指定区间内的成员
	ZRANGE = "ZRANGE"
	//	ZRANGEBYLEX key min max [LIMIT offset count] 通过字典区间返回有序集合的成员
	ZRANGEBYLEX = "ZRANGEBYLEX"
	//	ZRANGEBYSCORE key min max [WITHSCORES] [LIMIT] 通过分数返回有序集合指定区间内的成员
	ZRANGEBYSCORE = "ZRANGEBYSCORE"
	//ZRANK key member 返回有序集合中指定成员的索引
	ZRANK = "ZRANK"
	//ZREM key member [member ...] 移除有序集合中的一个或多个成员
	ZREM = "ZREM"
	//	ZREMRANGEBYLEX key min max 移除有序集合中给定的字典区间的所有成员
	ZREMRANGEBYLEX = "ZREMRANGEBYLEX"
	//	ZREMRANGEBYRANK key start stop 移除有序集合中给定的排名区间的所有成员
	ZREMRANGEBYRANK = "ZREMRANGEBYRANK"
	//ZREMRANGEBYSCORE key min max 移除有序集合中给定的分数区间的所有成员
	ZREMRANGEBYSCORE = "ZREMRANGEBYSCORE"
	//ZREVRANGE key start stop [WITHSCORES] 返回有序集中指定区间内的成员，通过索引，分数从高到底
	ZREVRANGE = "ZREVRANGE"
	//ZREVRANGEBYSCORE key max min [WITHSCORES] 返回有序集中指定分数区间内的成员，分数从高到低排序
	ZREVRANGEBYSCORE = "ZREVRANGEBYSCORE"
	//ZREVRANK key member 返回有序集合中指定成员的排名，有序集成员按分数值递减(从大到小)排序
	ZREVRANK = "ZREVRANK"
	//ZSCORE key member 返回有序集中，成员的分数值
	ZSCORE = "ZSCORE"
	//	ZUNIONSTORE destination numkeys key [key ...] 计算给定的一个或多个有序集的并集，并存储在新的 key 中
	ZUNIONSTORE = "ZUNIONSTORE"
	//	ZSCAN key cursor [MATCH pattern] [COUNT count] 迭代有序集合中的元素（包括元素成员和元素分值）
	ZSCAN = "ZSCAN"
	//SUBSCRIBE  发布订阅
	//PSUBSCRIBE pattern [pattern ...] 订阅一个或多个符合给定模式的频道
	PSUBSCRIBE = "PSUBSCRIBE"
	//PUBSUB subcommand [argument [argument ...]] 查看订阅与发布系统状态
	PUBSUB = "PUBSUB"
	//PUBLISH channel message 将信息发送到指定的频道
	PUBLISH = "PUBLISH"
	//PUNSUBSCRIBE [pattern [pattern ...]] 退订所有给定模式的频道
	PUNSUBSCRIBE = "PUNSUBSCRIBE"
	//SUBSCRIBE channel [channel ...] 订阅给定的一个或多个频道的信息
	SUBSCRIBE = "SUBSCRIBE"
	//UNSUBSCRIBE [channel [channel ...]] 指退订给定的频道
	UNSUBSCRIBE = "UNSUBSCRIBE"

	//事务
	//DISCARD 取消事务，放弃执行事务块内的所有命令
	DISCARD = "DISCARD"
	//EXEC 执行所有事务块内的命令
	EXEC = "EXEC"
	//MULTI 标记一个事务块的开始
	MULTI = "MULTI"
	//	UNWATCH 取消 WATCH 命令对所有 key 的监视
	UNWATCH = "UNWATCH"
	//WATCH key [key ...] 监视一个(或多个) key ，如果在事务执行之前这个(或这些) key 被其他命令所改动，那么事务将被打断
	WATCH = "WATCH"
	//连接
	//AUTH password 验证密码是否正确
	AUTH = "AUTH"
	//ECHO message 打印字符串
	ECHO = "ECHO"
	//PING 查看服务是否运行
	PING = "PING"
	//	QUIT 关闭当前连接
	QUIT = "QUIT"
	//SELECT index 切换到指定的数据库
	SELECT = "SELECT"
	//服务器命令
	//BGREWRITEAOF 异步执行一个 AOF（AppendOnly File） 文件重写操作
	BGREWRITEAOF = "BGREWRITEAOF"
	//BGSAVE 在后台异步保存当前数据库的数据到磁盘
	BGSAVE = "BGSAVE"
	//CLIENT KILL [ip:port] [ID client-id] 关闭客户端连接
	CLIENTKILL = "CLIENT KILL"
	//CLIENT LIST 获取连接到服务器的客户端连接列表
	CLIENTLIST = "CLIENT LIST"
	//CLIENT GETNAME 获取连接的名称
	CLIENTGETNAME = "CLIENT GETNAME"
	//CLIENT PAUSE timeout 在指定时间内终止运行来自客户端的命令
	CLIENTPAUSE = "CLIENT PAUSE"
	//	CLUSTER SLOTS 获取集群节点的映射数组
	CLUSTERSLOTS = "CLUSTER SLOTS"
	//COMMAND 获取 Redis 命令详情数组
	COMMAND = "COMMAND"
	//	TIME 返回当前服务器时间
	TIME = "TIME"
	//CONFIG REWRITE 对启动 Redis 服务器时所指定的 redis.conf 配置文件进行改写
	CONFIGREWRITE = "CONFIG REWRITE"
	//CONFIG SET parameter value 修改 redis 配置参数，无需重启
	CONFIGSET = "CONFIG SET"
	//CONFIG RESETSTAT 重置 INFO 命令中的某些统计数据
	CONFIGRESETSTAT = "CONFIG RESETSTAT"
	//DBSIZE 返回当前数据库的 key 的数量
	DBSIZE = "DBSIZE"
	//	DEBUG SEGFAULT 让 Redis 服务崩溃
	DEBUGSEGFAULT = "DEBUG SEGFAULT"
	//	FLUSHALL 删除所有数据库的所有key
	FLUSHALL = "FLUSHALL"
	//FLUSHDB 删除当前数据库的所有key
	FLUSHDB = "FLUSHDB"
	//INFO [section] 获取 Redis 服务器的各种信息和统计数值
	INFO = "INFO"
	//LASTSAVE 返回最近一次 Redis 成功将数据保存到磁盘上的时间，以 UNIX 时间戳格式表示
	LASTSAVE = "LASTSAVE"
	//MONITOR 实时打印出 Redis 服务器接收到的命令，调试用
	MONITOR = "MONITOR"
	//ROLE 返回主从实例所属的角色
	ROLE = "ROLE"
	//SAVE 同步保存数据到硬盘
	SAVE = "SAVE"
	//SHUTDOWN [NOSAVE] [SAVE] 异步保存数据到硬盘，并关闭服务器
	SHUTDOWN = "SHUTDOWN"
	//SLAVEOF host port 将当前服务器转变为指定服务器的从属服务器(slave server)
	SLAVEOF = "SLAVEOF"
	//SLOWLOG subcommand [argument] 管理 redis 的慢日志
	SLOWLOG = "SLOWLOG"
	//SYNC 用于复制功能(replication)的内部命令
	SYNC = "SYNC"
)

func connect() (redis.Conn, error) {
	con, err := redis.Dial("tcp", "127.0.0.1:6379",
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

//MyRedisErr err
type MyRedisErr struct {
	err string
}

func (redis *MyRedisErr) Error() string {
	return redis.err
}

//RedisSyncLock 分布式锁实现
type RedisSyncLock struct {
	Lockkey string
	Expr    int64
}

//Lock 锁
func (lock *RedisSyncLock) Lock(cnn redis.Conn) (bool, error) {
	if cnn == nil {
		return false, &MyRedisErr{err: "cnn is nil"}
	}
	for i := 1; i < 500; i++ {
		//过期时间
		tm := time.Now().Unix() + lock.Expr
		//设置锁
		result, err := cnn.Do(SETNX, lock.Lockkey, strconv.FormatInt(tm, 0))
		//未知错误
		if err != nil {
			return false, err
		}
		//设置成功
		if result != nil {
			return true, nil
		}
		//获取旧的值
		result, err = cnn.Do(GET, lock.Lockkey)
		if err != nil {
			return false, err
		}
		//如果获取到的是空值那么就设置成0
		if result == nil {
			result = "0"
		}
		oldvalue, _ := strconv.Atoi(result.(string))
		//检查时间是否是已经过期
		if uint32(oldvalue) <= uint32(time.Now().Unix()) {
			//如果已经过期直接getset
			result, err = cnn.Do(GETSET, lock.Lockkey, strconv.FormatInt(tm, 0))
			newvalue, _ := strconv.Atoi(result.(string))
			//设置成功
			if oldvalue == newvalue {
				return true, nil
			}
		}
		//不成功休息下继续轮询
		select {
		//休息10ms
		case <-time.After(time.Millisecond * 10):
			continue
		}
	}
	return false, nil
}

/*
	理论上有可能会导致释放掉不是我自己加的锁,为了避免这种情况,
	再加锁的时候过期时间尽量设置的比较长一点,尽量不要产生超时被其它锁给抢占的情况
*/
//Unlock 释放锁
func (lock *RedisSyncLock) Unlock(cnn redis.Conn) {
	cnn.Do(DEL, lock.Lockkey)
}
