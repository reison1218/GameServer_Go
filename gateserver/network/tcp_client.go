package network

import (
	"gateserver/config_helper"
	"gateserver/log"
	"net"
	"tools/tcp"
)

var NetClient *ClientHandler

type ClientHandler struct {
	conn *net.TCPConn
}

func (client *ClientHandler) OnOpen(conn *net.TCPConn) {
	client.conn = conn
}

func (client *ClientHandler) OnClose() {
	log.Logs.Error("断开链接")
}
func (client *ClientHandler) OnMessage(bytes []byte) {

}

func InitClient() {
	NetClient = &ClientHandler{}
	value := config_helper.Configs.Configs["game_port"]
	address := value.String()
	tcp.InitTcpClient(address, NetClient)
}
