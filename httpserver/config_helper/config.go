package config_helper

import (
	"os"
	"tools"
	"tools/utils"
)

var Configs utils.Conf

func InitConfig() utils.Conf {
	dirPath, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	//配置文件路径
	configFile := dirPath + "/config/config.conf"
	return tools.InitLog(configFile)
}
