package mgr

import (
	"httpserver/entity"
	"httpserver/log"
	"httpserver/mysql_pool"
)

var SIM *ServerInfoMgr = &ServerInfoMgr{}

type ServerInfoMgr struct {
	ServerMap map[int]*entity.ServerInfo
}

func queryAll() {
	serverMap, err := entity.QueryAll(mysql_pool.MysqlPool)
	if err != nil {
		log.Logs.Error(err.Error())
		return
	}

	for _, si := range serverMap {
		var v entity.ServerInfo = si.(entity.ServerInfo)
		SIM.ServerMap[v.ServerId] = &v
	}

}

func (sim *ServerInfoMgr) GetAll() []*entity.ServerInfo {
	// []*entity.ServerInfo
	var array []*entity.ServerInfo
	var index = 0
	for _, v := range sim.ServerMap {
		array[index] = v
		index += 1
	}
	return array
}

func InitMgr() {
	SIM.ServerMap = make(map[int]*entity.ServerInfo)
	queryAll()
}
