package main

import (
	"httpserver/config_helper"
	"httpserver/log"
	"httpserver/mgr"
	"httpserver/mysql_pool"
	"httpserver/network"
	"time"
	"tools/utils"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	startTime := time.Now()
	//初始化配置文件
	initConfig()
	//初始化日志
	initLog()
	//初始化db
	initMysql()
	//初始化管理器
	initMgr()
	//初始化tcpclient
	initTcpClient()
	//初始化httpserver
	initHttpServer()
	endTime := time.Now()
	res := endTime.Sub(startTime)
	log.Logs.Info("httpserver init success!take time:" + res.String())
	select {}
}

func initMgr() {
	mgr.InitMgr()
}

func initMysql() {
	mysql_pool.InitMysql()
}

func initHttpServer() {
	network.InitHttpServer()
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

func initTcpClient() {
	go network.InitClient()
}
