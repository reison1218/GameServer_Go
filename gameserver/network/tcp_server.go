package network

import (
	"errors"
	"fmt"
	"gameserver/config_helper"
	"gameserver/entity"
	"gameserver/log"
	"gameserver/mgr"
	"gameserver/redis_client"
	"strconv"
	"sync"
	"tools/cmds"
	"tools/protos"
	"tools/tcp"
	"tools/utils"

	jsoniter "github.com/json-iterator/go"
	"google.golang.org/protobuf/proto"
)

type ServerHandler struct {
	Token uint32
	Conn  *tcp.TcpSession
	lock  *sync.RWMutex
}

func (serverHandler *ServerHandler) OnMessage(bytes []byte) {
	log.Logs.Info("there is mess from client!size:" + strconv.Itoa(len(bytes)))
	packetList, error := utils.BuildArrayFromClient(bytes)
	if error != nil {
		log.Logs.Error(error.Error())
		return
	}

	//从屁股拿
	for i := packetList.Back(); i != nil; i = i.Next() {
		packet := i.Value.(utils.Packet)
		serverHandler.handler_binary(packet)
	}
}

func (serverHandler *ServerHandler) Clone() tcp.ServerHandlerTrigger {
	return &ServerHandler{}
}

func (serverHandler *ServerHandler) OnClose() {
	log.Logs.Info("client disconnect!")
}

func (serverHandler *ServerHandler) OnOpen(tcpConn *tcp.TcpSession) {
	log.Logs.Info("there is new client connect!")
	serverHandler.Token = tcpConn.Token
	serverHandler.Conn = tcpConn
}

func (serverHandler *ServerHandler) SetToken(token uint32) {
	serverHandler.Token = token
}

func (serverHandler *ServerHandler) handler_binary(packet utils.Packet) {
	//处理来自gateserver的消息
	cmd := packet.Des.Cmd
	if cmd != cmds.LOGIN && cmd != cmds.UNLOADUSER && len(packet.GetData()) == 0 {
		log.Logs.Error("packet bytes is null!cmd:" + strconv.Itoa(int(cmd)))
		return
	}
	handler_message(&packet)
}

func handler_message(packet *utils.Packet) {
	cmd := packet.Des.Cmd
	if cmd == cmds.LOGIN {
		protoReq := protos.C_USER_LOGIN{}
		err := proto.Unmarshal(packet.GetData(), &protoReq)
		if err != nil {
			log.Logs.Error(err.Error())
			return
		}
		Login(mgr.GM, packet)
	} else {
		fn := mgr.GM.CmdMap[int(cmd)]
		if fn == nil {
			str := fmt.Sprintf("can not find func for cmd:%d", cmd)
			log.Logs.Error(str)
			return
		}
		fn(mgr.GM, packet)
	}
}

func Login(gm *mgr.GameMgr, packet *utils.Packet) {
	userId := packet.Des.UserId
	userData := gm.Users[int(userId)]
	//如果内存没有数据，则从数据库里面找
	if userData == nil {
		res, err := initUserData(int(userId))
		if err != nil {
			log.Logs.Error(err.Error())
			return
		}
		gm.Users[int(userId)] = res
	}
	userData = gm.Users[int(userId)]
	userData.UpdateLogin()
	//返回客户端
	protoRes := protos.S_USER_LOGIN{}
	protoRes.IsSucc = true
	bytes, _ := proto.Marshal(&protoRes)
	gm.TcpConn.Write(bytes)
}

func initUserData(userId int) (*entity.UserData, error) {
	var userData *entity.UserData
	var err error
	var jsonData string
	//从redis拿
	jsonData, err = getUserFromRedis(userId)
	if err != nil {
		return nil, err
	}
	if jsonData == "" {
		str := fmt.Sprintf("redis has no data for userId:%d", userId)
		return nil, errors.New(str)
	}
	//从数据库查
	userData, err = entity.QueryUserData(userId)
	if err != nil {
		return nil, err
	}
	//数据库没有就新建数据
	if userData == nil {
		var jsonMap map[string]jsoniter.Number

		err := jsoniter.Unmarshal([]byte(jsonData), &jsonMap)
		if err != nil {
			return nil, err
		}
		jsonValue := jsonMap["nick_name"]
		if jsonMap == nil || jsonValue == "" {
			return nil, errors.New("jsonValue has no data for nick_name")
		}
		userData = &entity.UserData{}
		userData.User.UserId = userId
		userData.User.NickName = jsonValue.String()
		userData.User.Grade = 1

		//初始化角色
		cters := &entity.Characters{}
		cters.UserId = userId
		userData.Cters = cters

		//初始化灵魂头像
		soul := &entity.Soul{}
		soul.UserId = userId
		userData.Soul = soul

		//初始化头像框
		gradeFrame := &entity.GradeFrame{}
		gradeFrame.UserId = userId
		userData.GradeFrame = gradeFrame

		//持久化道数据库
		go entity.InsertUserData(userData)
	}
	return userData, nil
}

func getUserFromRedis(userId int) (string, error) {
	redis_client.RedisClient.Do("select", 0)
	res, err := redis_client.RedisClient.Do("hget", "uid_2_pid", strconv.Itoa(userId))
	if err != nil {
		return "", err
	}
	pid := res.(string)
	jsonRes, err := redis_client.RedisClient.Do("hget", "users", pid)
	if err != nil {
		return "", err
	}
	jsonStr := jsonRes.(string)
	return jsonStr, nil
}

func InitServer() {
	value := config_helper.Configs.Configs["tcp_port"]
	address := value.String()
	sh := ServerHandler{}
	tcp.InitTcpServer(address, &sh)
}
