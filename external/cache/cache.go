package cache

import (
	"fmt"

	"github.com/gomodule/redigo/redis"
	"github.com/jwilyandi19/simple-product/helper"
)

type Cache struct {
	Redis *redis.Pool
	TTL   int
}

func InitCacheConnection(conf helper.RedisConfig) Cache {
	addr := fmt.Sprintf("%s:%d", conf.Server, conf.Port)
	redis := &redis.Pool{
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", addr, redis.DialPassword(conf.Password))
		},
	}
	return Cache{
		Redis: redis,
		TTL:   conf.TTL,
	}
}
