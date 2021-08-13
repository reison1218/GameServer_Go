package redis_helper

import (
	"activitycenter/config_helper"
	"fmt"
	"log"

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

func Init() {
	reidsAddress := config_helper.Configuration.Configs["redis_add"]
	reidsPass := config_helper.Configuration.Configs["redis_pass"]
	reidsClient, err := redis.DialURL(reidsAddress.String())
	if err != nil {
		panic(err)
	}
	_, err = reidsClient.Do("auth", reidsPass.String())
	if err != nil {
		panic(err)
	}
	RedisGlobalHelper = reidsClient
	log.Println("redis_client init success!")
}
