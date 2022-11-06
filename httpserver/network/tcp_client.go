package network

import (
	"httpserver/config_helper"
	"httpserver/log"
	"httpserver/mgr"
	"httpserver/protos"
	"net"
	"time"
	"tools/tcp"
	"tools/utils"

	"github.com/golang/protobuf/proto"
)

var NetClient *ClientHandler

type ClientHandler struct {
	conn *net.TCPConn
}

func (client *ClientHandler) OnOpen(conn *net.TCPConn) {
	client.conn = conn

	//发送握手协议
	client.sendHandShake(protos.CommandType_CommandType_HandShake, protos.CommandType_CommandType_HandShake)
}

func (client *ClientHandler) OnClose() {
	log.Logs.Error("断开链接")
}

func (client *ClientHandler) OnMessage(bytes []byte) {
	bb := utils.NewByteBuffer()
	bb.WriteBytes(bytes)

	bodyLen, bodyErr := bb.ReadUint16()
	if bodyErr != nil {
		log.Logs.Error(bodyErr.Error())
		return
	}
	headerLen, headerErr := bb.ReadUint16()
	if headerErr != nil {
		log.Logs.Error(headerErr.Error())
		return
	}
	headerBytes, headerBytesErr := bb.ReadBytes(int(headerLen))
	if headerBytesErr != nil {
		log.Logs.Error(headerBytesErr.Error())
		return
	}

	var _, err = bb.ReadBytes(int(bodyLen))
	if err != nil {
		log.Logs.Error(err.Error())
		return
	}
	th := protos.TMSG_HEADER{}
	errr := proto.Unmarshal(headerBytes, &th)
	if errr != nil {
		log.Logs.Error(errr.Error())
		return
	}
	//如果是握手协议
	if *th.Command.Enum() == protos.CommandType_CommandType_HandShakeResponse {
		log.Logs.Info("handshake with socket-server success!")
		go client.keepLive()
	}
	//如果是心跳协议
	if *th.Command.Enum() == protos.CommandType_CommandType_KeepaliveResponse {
		log.Logs.Info("收到心跳返回")
	}
	//如果是邮件协议
	if *th.Command.Enum() == protos.CommandType_CommandType_OperationEmail {
		log.Logs.Info("收到邮件协议")

		_, bodyBytesErr := bb.ReadBytes(int(bodyLen - headerLen))
		if bodyBytesErr != nil {
			log.Logs.Error(bodyBytesErr.Error())
			return
		}
	}
}

// 循环发送心跳包
func (client *ClientHandler) keepLive() {
	for {
		time.Sleep(time.Millisecond * 3000)
		_, err := client.send(protos.CommandType_CommandType_Keepalive, protos.CommandType_CommandType_Keepalive, make([]byte, 0))
		if err != nil {
			log.Logs.Error(err.Error())
			return
		}
	}
}

// 发送握手协议
func (client *ClientHandler) sendHandShake(cmd protos.CommandType, subCmd protos.CommandType) {
	//拿到游戏id
	var configId = config_helper.Configs.Configs["game_id"]
	id, err := configId.Int64()
	if err != nil {
		log.Logs.Error("找不到game_id!" + err.Error())
		return
	}
	//游戏id
	var gameId = uint32(id)

	//服务器列表id
	var worldIds = make([]uint32, len(mgr.SIM.ServerMap))

	index := 0
	for id, _ := range mgr.SIM.ServerMap {
		worldIds[index] = uint32(id)
		index += 1
	}

	//封装proto
	var body = protos.TMSG_HANDSHAKE_REQ{}
	body.GameID = &gameId
	body.WorldIDs = worldIds
	//decode
	bytes, err := proto.Marshal(&body)
	if err != nil {
		log.Logs.Error(err.Error())
		return
	}
	//发送
	client.send(cmd, subCmd, bytes)
}

// 发送命令
func (client *ClientHandler) send(cmd protos.CommandType, subCmd protos.CommandType, bytes []byte) (int, error) {
	var gameId uint32 = 2141
	var worldId uint32 = 2141
	var worldIds = make([]uint32, 1)
	worldIds[0] = gameId

	var header = protos.TMSG_HEADER{}
	header.GameID = &gameId
	header.Command = cmd.Enum()
	header.WorldID = &worldId

	headerBytes, err := proto.Marshal(&header)
	if err != nil {
		log.Logs.Error(err.Error())
		return 0, nil
	}

	headBytesLen := uint16(len(headerBytes))

	bodyBytesLen := uint16(len(bytes))

	var bodyLen uint16 = uint16(2) + headBytesLen + bodyBytesLen

	var bb = &utils.ByteBuffer{}
	bb.WriteUint16(bodyLen)
	bb.WriteUint16(headBytesLen)
	bb.WriteBytes(headerBytes)
	bb.WriteBytes(bytes)
	return client.conn.Write(bb.GetAllBytes())
}

// 初始化tcp客户端
func InitClient() {
	NetClient = &ClientHandler{}
	value := config_helper.Configs.Configs["socker_server"]
	address := value.String()
	tcp.InitTcpClient(address, NetClient)
}
