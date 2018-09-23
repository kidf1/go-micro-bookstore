package redis

import (
	"time"

	"github.com/gomodule/redigo/redis"
)

// Host redis default
var (
	Host = "127.0.0.1:6379"
	Pool redis.Pool
)

// NewPool redis
func NewPool() *redis.Pool {

	return &redis.Pool{

		MaxIdle:     3,
		IdleTimeout: 240 * time.Second,

		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", Host)
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
}
