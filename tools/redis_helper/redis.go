package redis_helper

import (
	"fmt"
	"tools/utils"

	"github.com/gomodule/redigo/redis"
)

func RedisTest(redisAddr string) {
	reidsClient, err := redis.DialURL("redis://" + redisAddr)
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

func InitRedis(redisAddr string, redisPass string) redis.Conn {
	reidsClient, err := redis.DialURL(redisAddr)
	if err != nil {
		panic(err)
	}
	_, err = reidsClient.Do("auth", redisPass)
	if err != nil {
		panic(err)
	}
	utils.ZapLogger.Info("redis_client init success!")
	return reidsClient
}
