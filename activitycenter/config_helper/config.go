package config_helper

import (
	"io/ioutil"
	"log"
	"os"

	jsoniter "github.com/json-iterator/go"
)

var Configuration Conf

type Conf struct {
	Configs map[string]jsoniter.Number
}

func newConf() Conf {
	conf := Conf{make(map[string]jsoniter.Number)}
	return conf
}

func Init() {
	Configuration = newConf()
	dirPath, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	//配置文件路径
	configFile := dirPath + "/config/config.conf"
	bytes, err := ioutil.ReadFile(configFile)
	if err != nil {
		panic(err)
	}

	errJson := jsoniter.Unmarshal(bytes, &Configuration.Configs)
	if errJson != nil {
		panic(errJson)
	}
	log.Println("config init success!")
}
