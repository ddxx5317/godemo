package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

var pool *redis.Pool

func init() {
	pool = &redis.Pool{
		MaxIdle:     8,
		MaxActive:   0,
		IdleTimeout: 100,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", "localhost:6379")
		},
	}
}

func main() {
	conn := pool.Get()
	defer conn.Close()

	_, err1 := conn.Do("HMSet", "user1", "name", "beijing", "address", "beijing")
	if err1 != nil {
		fmt.Println("HMSet err=", err1)
		return
	}
	_, err3 := conn.Do("HMSet", "user2", "name", "wuhan", "address", "wuhan")
	if err3 != nil {
		fmt.Println("HMSet err=", err3)
		return
	}

	//向redis读取数据，返回的r是个空接口
	r, err2 := redis.Strings(conn.Do("HMGet", "user1", "name", "address"))
	if err2 != nil {
		fmt.Println("HMGet err=", err2)
		return
	}
	for i, v := range r {
		fmt.Printf("r[%d]=%v\n", i, v)
	}
}