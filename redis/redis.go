package redis

import (
	"github.com/gomodule/redigo/redis"
)

var Cache redis.Conn


func InitCache(host string) {
	// init the redis connection
	conn,err := redis.Dial("tcp",host)
	if err != nil {
		panic(err.Error()+" asd")
	}
	// assign the connection to cache variable
	Cache = conn
}
