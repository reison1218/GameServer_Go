package redis_helper

import (
	"fmt"

	"github.com/gomodule/redigo/redis"
)

var RedisGlobalHelper redis.Conn

func RedisTest() {
	reidsClient, err := redis.DialURL("redis://127.0.0.1:6379")
	if err != nil {
		panic(err)
	}
	_, err = reidsClient.Do("auth", "reison")
	if err != nil {
		panic(err)
	}
	res, err := reidsClient.Do("hget", "users", "test")
	if err != nil {
		panic(err)
	}

	fmt.Println(string(res.([]byte)))
	defer reidsClient.Close()
}

func Init() redis.Conn {
	reidsClient, err := redis.DialURL("redis://127.0.0.1:6379")
	if err != nil {
		panic(err)
	}
	_, err = reidsClient.Do("auth", "reison")
	if err != nil {
		panic(err)
	}
	return reidsClient
}
