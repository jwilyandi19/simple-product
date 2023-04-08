package cache

import (
	"github.com/gomodule/redigo/redis"
	"github.com/jwilyandi19/simple-product/helper"
)

type Cache struct {
	Redis *redis.Pool
}

func InitCacheConnection(conf helper.RedisConfig) Cache {
	redis := &redis.Pool{
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", conf.Server, redis.DialPassword(conf.Password))
		},
	}
	return Cache{
		Redis: redis,
	}
}
