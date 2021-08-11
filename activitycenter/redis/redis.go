package redis

import (
	"fmt"

	"github.com/gomodule/redigo/redis"
)

func RedisTest() {
	reidsClient, err := redis.DialURL("redis://127.0.0.1:6379")
	if err != nil {
		panic(err)
		return
	}
	_, err = reidsClient.Do("auth", "reison")
	if err != nil {
		panic(err)
		return
	}
	res, err := reidsClient.Do("hget", "users", "test")
	if err != nil {
		panic(err)
		return
	}

	fmt.Println(string(res.([]byte)))
	defer reidsClient.Close()
}
