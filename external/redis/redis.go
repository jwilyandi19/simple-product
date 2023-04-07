package redis

import (
	"github.com/gomodule/redigo/redis"
	"github.com/jwilyandi19/simple-product/helper"
)

func InitRedisConnection(conf helper.RedisConfig) *redis.Pool {
	return &redis.Pool{
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", conf.Server, redis.DialPassword(conf.Password))
		},
	}
}
