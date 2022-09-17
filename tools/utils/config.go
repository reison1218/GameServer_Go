package utils

import (
	"io/ioutil"

	jsoniter "github.com/json-iterator/go"
)

type Conf struct {
	Configs map[string]jsoniter.Number
}

func newConf() Conf {
	conf := Conf{make(map[string]jsoniter.Number)}
	return conf
}

func Init(path string) Conf {
	Configuration := newConf()

	//配置文件路径
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}

	errJson := jsoniter.Unmarshal(bytes, &Configuration.Configs)
	if errJson != nil {
		panic(errJson)
	}
	return Configuration
}
