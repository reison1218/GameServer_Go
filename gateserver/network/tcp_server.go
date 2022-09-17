package network

import (
	"errors"
	"gateserver/config_helper"
	"gateserver/log"
	"gateserver/mgr"
	"gateserver/redis_client"
	"strconv"
	"sync"
	"time"
	"tools/cmds"
	"tools/protos"
	"tools/tcp"
	"tools/utils"

	"github.com/golang/protobuf/proto"
)

type ServerHandler struct {
	Token uint32
	Conn  *tcp.TcpSession
	lock  *sync.RWMutex
}

func (sh *ServerHandler) Write(bytes []byte) {
	sh.Conn.Conn.Write(bytes)
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

func (serverHandler *ServerHandler) handler_binary(packet utils.Packet) {
	token := serverHandler.Token
	serverHandler.lock.RLock()
	userId := mgr.CM.Channel[token]

	cmd := packet.Des.Cmd
	//如果没有缓存，并且又不是登陆，则视作违法消息
	if userId == 0 && cmd != cmds.LOGIN {
		log.Logs.Error("this player is not login and cmd != Login!cmd:" + strconv.Itoa(int(cmd)) + ",token:" + strconv.Itoa(int(token)))
		return
	}
	//如果等于登陆，直接转发给游戏服务器
	if cmd == cmds.LOGIN {

		loginProto := &protos.C_USER_LOGIN{}
		err := proto.Unmarshal(packet.GetData(), loginProto)
		if err != nil {
			log.Logs.Error(err.Error())
			return
		}
		platformValue := loginProto.PlatformValue
		registerPlatform := loginProto.RegisterPlatform
		protoUserId := loginProto.UserId
		//处理登陆
		res, err := handlerLogin(&mgr.CM, registerPlatform, platformValue, protoUserId)
		if err != nil {
			return
		}
		userId = uint32(res)
	}
	//设值
	packet.SetUserId(userId)
	//如果等于心跳
	if cmd == cmds.HEARTBEAT {
		hbProto := &protos.HEART_BEAT{}
		err := proto.Unmarshal(packet.GetData(), hbProto)
		if err != nil {
			log.Logs.Error(err.Error())
			return
		}
		hbProto.SysTime = uint64(time.Now().UnixMilli())
		gateUser := mgr.CM.UserChannel[userId]
		bytes, _ := proto.Marshal(hbProto)
		gateUser.SS.Write(bytes)
		log.Logs.Info("回客户端消息,user_id:" + strconv.Itoa(int(userId)) + ",cmd:" + strconv.Itoa(int(cmd)))
		return
	}
	//如果都不是，就执行转发函数
	arrangePacket(&mgr.CM, &packet)
}

/*消息分发*/
func arrangePacket(cm *mgr.ChannelMgr, packet *utils.Packet) {
	cmd := packet.Des.Cmd
	if cmd <= cmds.GAME_MAX || cmd >= cmds.GAME_MIN {
		WriteToGame(cm, packet)
	}
}

/*写给服务器*/
func WriteToGame(cm *mgr.ChannelMgr, packet *utils.Packet) {
	if cm.GameClientChannel == nil {
		log.Logs.Error("disconnect with game-server!")
		return
	}
	cm.GameClientChannel.Write(packet.BuildServerBytes())
}

/*处理登陆*/
func handlerLogin(cm *mgr.ChannelMgr, registerPlatform string, platformValue string, userId uint32) (int, error) {
	_, err := QueryPidFromRedis(userId)
	if err != nil {
		log.Logs.Error(err.Error())
		return 0, err
	}
	//检查内存里面有没有
	memRes := checkMemOnline(userId, cm)
	if memRes {
		log.Logs.Warn("发现重复登陆！T掉之前的tcp！userId:" + strconv.Itoa(int(userId)))
		return int(userId), errors.New("重复登陆！")
	}
	return int(userId), nil
}

func checkMemOnline(userId uint32, cm *mgr.ChannelMgr) bool {
	gateUser := cm.UserChannel[userId]
	return gateUser == nil
}

func QueryPidFromRedis(userId uint32) (string, error) {
	userIdStr := strconv.Itoa(int(userId))
	res, err := redis_client.RedisClient.Do("hget", "0", "uid_2_pid", userIdStr)
	if err != nil {
		log.Logs.Error(err.Error())
		return "", err
	}
	return res.(string), nil
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

func InitServer() {
	value := config_helper.Configs.Configs["tcp_port"]
	address := value.String()
	sh := ServerHandler{}
	tcp.InitTcpServer(address, &sh)
}
