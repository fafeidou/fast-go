package gredis

import (
	"encoding/json"
	"fast-go/conf"
	"time"

	"github.com/gomodule/redigo/redis"
)

var RedisConn *redis.Pool

// Setup Initialize the Redis instance
func Setup() error {
	RedisConn = &redis.Pool{
		MaxIdle:    conf.App.Redis.MaxIdle,
		MaxActive:   conf.App.Redis.MaxActive,
		IdleTimeout: conf.App.Redis.IdleTimeout,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", conf.App.Redis.Host)
			if err != nil {
				return nil, err
			}
			if conf.App.Redis.Password != "" {
				if _, err := c.Do("AUTH", conf.App.Redis.Password); err != nil {
					c.Close()
					return nil, err
				}
			}
			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}

	return nil
}

// Set a key/value
func Set(key string, data interface{}, time int) error {
	conn := RedisConn.Get()
	defer conn.Close()

	value, err := json.Marshal(data)
	if err != nil {
		return err
	}

	_, err = conn.Do("SET", key, value)
	if err != nil {
		return err
	}

	_, err = conn.Do("EXPIRE", key, time)
	if err != nil {
		return err
	}

	return nil
}

// Exists check a key
func Exists(key string) bool {
	conn := RedisConn.Get()
	defer conn.Close()

	exists, err := redis.Bool(conn.Do("EXISTS", key))
	if err != nil {
		return false
	}

	return exists
}

// Get get a key
func Get(key string) ([]byte, error) {
	conn := RedisConn.Get()
	defer conn.Close()

	reply, err := redis.Bytes(conn.Do("GET", key))
	if err != nil {
		return nil, err
	}

	return reply, nil
}

// Delete delete a kye
func Delete(key string) (bool, error) {
	conn := RedisConn.Get()
	defer conn.Close()

	return redis.Bool(conn.Do("DEL", key))
}

// LikeDeletes batch delete
func LikeDeletes(key string) error {
	conn := RedisConn.Get()
	defer conn.Close()

	keys, err := redis.Strings(conn.Do("KEYS", "*"+key+"*"))
	if err != nil {
		return err
	}

	for _, key := range keys {
		_, err = Delete(key)
		if err != nil {
			return err
		}
	}

	return nil
}

// set expires为0时，表示永久性存储
func SetKey(key, value interface{}, expires int) error {
	rds := RedisConn.Get()
	defer rds.Close()
	if expires == 0 {
		_, err := rds.Do("SET", key, value)
		return err
	} else {
		_, err := rds.Do("SETEX", key, expires, value)
		return err
	}
}

// del
func DelKey(key string) error {
	rds := RedisConn.Get()
	defer rds.Close()
	_, err := rds.Do("DEL", key)
	return err
}

// lrange
func LRange(key string, start, stop int64) ([]string, error) {
	rds := RedisConn.Get()
	defer rds.Close()
	return redis.Strings(rds.Do("LRANGE", key, start, stop))
}

// lpop
func LPop(key string) (string, error) {
	rds := RedisConn.Get()
	defer rds.Close()
	return redis.String(rds.Do("LPOP", key))
}

// LPushAndTrimKey
func LPushAndTrimKey(key, value interface{}, size int64) error {
	rds := RedisConn.Get()
	defer rds.Close()
	rds.Send("MULTI")
	rds.Send("LPUSH", key, value)
	rds.Send("LTRIM", key, size-2*size, -1)
	_, err := rds.Do("EXEC")
	return err
}

// RPushAndTrimKey
func RPushAndTrimKey(key, value interface{}, size int64) error {
	rds := RedisConn.Get()
	defer rds.Close()
	rds.Send("MULTI")
	rds.Send("RPUSH", key, value)
	rds.Send("LTRIM", key, size-2*size, -1)
	_, err := rds.Do("EXEC")
	return err

}

// ExistsKey
func ExistsKey(key string) (bool, error) {
	rds := RedisConn.Get()
	defer rds.Close()
	return redis.Bool(rds.Do("EXISTS", key))
}

// ttl 返回剩余时间
func TTLKey(key string) (int64, error) {
	rds := RedisConn.Get()
	defer rds.Close()
	return redis.Int64(rds.Do("TTL", key))
}

// incr 自增
func Incr(key string) (int64, error) {
	rds := RedisConn.Get()
	defer rds.Close()
	return redis.Int64(rds.Do("INCR", key))
}

// Decr 自减
func Decr(key string) (int64, error) {
	rds := RedisConn.Get()
	defer rds.Close()
	return redis.Int64(rds.Do("DECR", key))
}

// mset 批量写入 rds.Do("MSET", "ket1", "value1", "key2","value2")
func MsetKey(key_value ...interface{}) error {
	rds := RedisConn.Get()
	defer rds.Close()
	_, err := rds.Do("MSET", key_value...)
	return err
}

// mget  批量读取 mget key1, key2, 返回map结构
func MgetKey(keys ...interface{}) map[interface{}]string {
	rds := RedisConn.Get()
	defer rds.Close()
	values, _ := redis.Strings(rds.Do("MGET", keys...))
	resultMap := map[interface{}]string{}
	keyLen := len(keys)
	for i := 0; i < keyLen; i++ {
		resultMap[keys[i]] = values[i]
	}
	return resultMap
}
