package redis

import (
	"github.com/gomodule/redigo/redis"
)

var Cache redis.Conn


func InitCache() {
	// init the redis connection
	conn,err := redis.DialURL("redis://localhost")
	if err != nil {
		panic(err)
	}
	// assign the connection to cache variable
	Cache = conn
}
