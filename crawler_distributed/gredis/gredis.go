package gredis

import (
	"github.com/gomodule/redigo/redis"
	"github.com/prometheus/common/log"
	"island/crawler_distributed/config"
	"time"
)

var RedisConn *redis.Pool

// Setup Initialize the Redis instance
func Setup() {
	RedisConn = &redis.Pool{
		MaxIdle:     config.MaxIdle,
		MaxActive:   config.MaxActive,
		IdleTimeout: config.IdleTimeout,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", config.RedisHost,
				redis.DialConnectTimeout(time.Duration(3000)*time.Millisecond))
			if err != nil {
				return nil, err
			}
			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}

	return
}


// Set redis hash表的使用
func SetBit(key string, field, value interface{}) {
	conn := RedisConn.Get()
	defer conn.Close()
	_, err := conn.Do("SETBIT", key, field, value)
	if err != nil {
		log.Error("pkg.gredis.SETBIT Do error is:", err)
		return
	}
	return
}

// GetBit redis hash表的使用
func GetBit(key string, field interface{}) (interface{}, error) {
	conn := RedisConn.Get()
	defer conn.Close()
	reply, err := conn.Do("GetBit", key, field)
	if err != nil {
		log.Error("pkg.gredis.HGet Do error is:", err)
		return nil, err
	}
	return reply, nil
}
