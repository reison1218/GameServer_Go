package main

import (
	"gateserver/config_helper"
	"gateserver/log"
	"gateserver/network"
	"gateserver/redis_client"
	"tools/utils"
)

type TestInterface interface {
	Haha(i int32)
	Clone() TestInterface
	GetI() int32
}

type Test struct {
	i int32
}

func (t *Test) Haha(i int32) {
	t.i = i
}

func (t *Test) Clone() TestInterface {
	return &Test{}
}

func (t *Test) GetI() int32 {
	return t.i
}

func aa(ti TestInterface) {
	res := ti.Clone()
	res.Haha(1)
	println(res.GetI())
	// var sh *Test = ti.(*Test)
	// res := sh.Clone()
	// var re Test = res.(Test)
	// re.Haha(1)
	// println(re.i)
}

func main() {
	//初始化配置文件
	initConfig()
	//初始化日志
	initLog()
	//初始化redis客户端
	initRedis()
	//初始化tcpclient
	initTcpClient()
	//初始化tcpserver
	initTcpServer()

}

func initRedis() {
	redis_client.InitReids()
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

func initTcpServer() {
	network.InitServer()
}

func initTcpClient() {
	go network.InitClient()
}
