package main

import (
	"gameserver/config_helper"
	"gameserver/entity"
	"gameserver/log"
	"gameserver/mysql_pool"
	"gameserver/network"
	"gameserver/redis_client"
	"strconv"
	"tools/utils"

	_ "github.com/go-sql-driver/mysql"
	jsoniter "github.com/json-iterator/go"
)

type TestInterface interface {
	K()
	Clone() interface{}
}

func main() {
	// 初始化配置
	initConfig()
	//初始化日志
	initLog()
	//初始化mysql
	initMysql()
	//初始化redis
	initRedis()
	//初始化tcp
	initTcpServer()
	test_redis()
	test_mysql()
	select {}
}

func test_redis() {
	res, err := redis_client.RedisClient.Do("hget", "uid_2_pid", strconv.Itoa(1011000001))
	if err != nil {
		panic(err)
	}
	if res == nil {
		println("reids nil")
	} else {
		println(res)
	}
}
func test_mysql() {
	userInfo := entity.UserInfo{}

	rows, err := mysql_pool.MysqlPool.Query("select * from t_u_player where user_id=?", 123)
	if err != nil {
		panic(err)
	}
	for rows.Next() {
		content := ""
		errrr := rows.Scan(&userInfo.UserId, &content)
		if errrr != nil {
			panic(errrr.Error())
		}
		jsonStr := content
		println("=================================:" + jsonStr)
		bytes := []byte(jsonStr)
		err := jsoniter.Unmarshal(bytes, &userInfo)
		if err != nil {
			panic(err)
		}
		userInfo.Ol = false
		ss, _ := jsoniter.Marshal(&userInfo)
		println(string(ss))

		res, err := mysql_pool.MysqlPool.Exec("update t_u_player set content=? where user_id=?", ss, 1011000001)
		if err != nil {
			panic(err)
		}
		println(res)

	}

}

func initMysql() {
	mysql_pool.InitMysql()
}

func initConfig() {
	config_helper.Configs = config_helper.InitConfig()
}
func initLog() {
	isDebug := config_helper.Configs.Configs["debug"]
	value, _ := isDebug.Int64()
	var logPath string
	if value == 0 {
		logPathValue := config_helper.Configs.Configs["log_path"]
		logPath = logPathValue.String()
	}
	appName := config_helper.Configs.Configs["app"]
	log.Logs = utils.InitLogger(logPath, appName.String(), value == 1)
}

func initRedis() {
	redis_client.InitReids()
}

func initTcpServer() {
	go network.InitServer()
}
