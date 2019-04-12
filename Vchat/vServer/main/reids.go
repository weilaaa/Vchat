package main

import (
	"github.com/garyburd/redigo/redis"
	"time"
)

var pool *redis.Pool

// init function of pool
func initPool(address string, maxIdle, maxActive int, idleTimeout time.Duration) {

	pool = &redis.Pool{
		MaxIdle:     maxIdle,
		MaxActive:   maxActive,
		IdleTimeout: idleTimeout,
		Dial: func() (conn redis.Conn, e error) {
			return redis.Dial("tcp", address)
		},
	}
}
