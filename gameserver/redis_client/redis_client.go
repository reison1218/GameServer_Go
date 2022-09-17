package redis_client

import (
	"gameserver/config_helper"
	"tools/redis_helper"

	"github.com/gomodule/redigo/redis"
)

var RedisClient redis.Conn

func InitReids() {
	valueAddress := config_helper.Configs.Configs["redis_add"]
	address := valueAddress.String()
	valuePass := config_helper.Configs.Configs["redis_pass"]
	pass := valuePass.String()
	RedisClient = redis_helper.InitRedis(address, pass)
}

func GetUserFromRedis(userId int) {}
