package main

import (
	"time"

	"github.com/gomodule/redigo/redis"
)

var pool *redis.Pool

func initPool(address string, maxIdle, maxActive int, idleTimeOut time.Duration) {
	pool = &redis.Pool{
		MaxIdle:     maxIdle,
		MaxActive:   maxActive, //最大连接数， 0表示没限制
		IdleTimeout: idleTimeOut,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", address)
		},
	}
}
