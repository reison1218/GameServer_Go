package entity

import (
	"database/sql"
	"fmt"
)

// 服务器信息结构体
type ServerInfo struct {
	ServerId       int    //服务器id
	Name           string //服务器名字
	Ws             string //ws连接
	OpenTime       string //开服时间
	RegisterState  int    //0超过N天不可注册1可注册
	State          int    //0无1爆2新
	Letter         int    //0正常1强行推荐
	TargetServerId int    //目标服id
	MergeTimes     int    //合服次数
	Type           string //类型：0开发 1内测 2正式
}

func QueryAll(sqlCon *sql.DB) (map[int]interface{}, error) {
	var table = "server_list"
	var sqlStr string
	var err error
	sqlStr = fmt.Sprintf("select * from %s ", table)

	rows, err := sqlCon.Query(sqlStr)
	if err != nil {
		return nil, err
	}
	defer func() {
		if rows != nil {
			rows.Close()
		}
	}()
	var errMess error

	var ServerId int
	var Name string        //服务器名字
	var Ws string          //ws连接
	var OpenTime string    //开服时间
	var RegisterState int  //0超过N天不可注册1可注册
	var State int          //0无1爆2新
	var Letter int         //0正常1强行推荐
	var TargetServerId int //目标服id
	var MergeTimes int     //合服次数
	var Type string        //类型：0开发 1内测 2正式

	var serverMap = make(map[int]interface{})
	for rows.Next() {
		errMess = rows.Scan(&ServerId, &Name, &Ws, &OpenTime, &RegisterState, &State, &Letter, &TargetServerId, &MergeTimes, &Type)
		if errMess != nil {
			return nil, errMess
		}

		//处理json
		// bytes := []byte(content)
		// res := en.Clone()
		// err = jsoniter.Unmarshal(bytes, res)
		// if err != nil {
		// 	panic(err)
		// }
		var si = ServerInfo{}
		si.ServerId = ServerId
		si.Name = Name
		si.Ws = Ws
		si.OpenTime = OpenTime
		si.RegisterState = RegisterState
		si.State = State
		si.Letter = Letter
		si.TargetServerId = TargetServerId
		si.MergeTimes = MergeTimes
		si.Type = Type
		serverMap[si.ServerId] = si
	}
	return serverMap, nil
}
