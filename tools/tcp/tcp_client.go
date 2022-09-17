package tcp

import (
	"net"
	"time"
	"tools/utils"
)

type ClientHandlerTrigger interface {
	OnClose()
	OnMessage(bytes []byte)
	OnOpen(conn *net.TCPConn)
}

func InitTcpClient(address string, clientHandler ClientHandlerTrigger) {
	// 绑定地址
	tcpAddr, err := net.ResolveTCPAddr("tcp", address)
	if err != nil {
		panic(err.Error())
	}
	//开始连接
	client, err := net.DialTCP("tcp", nil, tcpAddr)
	//如果失败了就一直连
	if err != nil {
		utils.ZapLogger.Error(err.Error())
		time.Sleep(time.Second * 5)
		InitTcpClient(address, clientHandler)
	}
	utils.ZapLogger.Info("connect '" + address + "' success!")
	//调用onopen函数
	clientHandler.OnOpen(client)
	//开始读取
	startClientRead(address, clientHandler, client)
}

//读取函数
func startClientRead(address string, clientHandler ClientHandlerTrigger, conn *net.TCPConn) {
	data := make([]byte, 512)
	var errStr string
	for {
		errStr = ""
		size, error := conn.Read(data)
		if size == 0 {
			errStr = address + " disconnected!"
		} else if error != nil {
			errStr = error.Error()
		}
		if errStr != "" {
			utils.ZapLogger.Error(errStr)
			clientHandler.OnClose()
			InitTcpClient(address, clientHandler)
			return
		}
		//正常读取数据
		clientHandler.OnMessage(data)
	}
}
