package tcp

import (
	"net"
	"sync"
	"tools/utils"
)

type TcpServer struct {
	tcpLister  *net.TCPListener
	tcpConnMap map[uint32]*net.TCPConn
}

func newTcpServer(tcpLister *net.TCPListener, tcpConnMap map[uint32]*net.TCPConn) TcpServer {
	var ts = new(TcpServer)
	ts.tcpLister = tcpLister
	ts.tcpConnMap = tcpConnMap
	return *ts
}

type ServerHandlerTrigger interface {
	Clone() ServerHandlerTrigger
	OnMessage(bytes []byte)
	OnClose()
	OnOpen(tcpConn *TcpSession)
	SetToken(token uint32)
}

type TcpSession struct {
	Conn  *net.TCPConn
	Token uint32
}

/*tcp服务器初始化*/
func InitTcpServer(address string, serverHandler ServerHandlerTrigger) {
	tcpAddr, err := net.ResolveTCPAddr("tcp", address) //创建 tcpAddr数据
	if err != nil {
		utils.ZapLogger.Error(err.Error())
		return
	}
	tcpListener, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		utils.ZapLogger.Error(err.Error())
		return
	}

	tcpServer := newTcpServer(tcpListener, make(map[uint32]*net.TCPConn))

	var lock sync.Mutex
	utils.ZapLogger.Info("start Listen:" + tcpServer.tcpLister.Addr().String())
	handlerMap := make(map[uint32]*ServerHandlerTrigger)
	for {
		tcpConn, err := tcpServer.tcpLister.AcceptTCP()
		if err != nil {
			utils.ZapLogger.Error(err.Error())
			continue
		}
		utils.ZapLogger.Info("there is new tcpClient coming:" + tcpConn.RemoteAddr().String())
		ID += 1
		ts := TcpSession{tcpConn, ID}
		sh := serverHandler.Clone()
		sh.OnOpen(&ts)
		sh.SetToken(ID)
		lock.Lock()
		handlerMap[ID] = &sh
		tcpServer.tcpConnMap[ID] = tcpConn
		lock.Unlock()
		go startRead(ID, tcpServer.tcpConnMap, handlerMap, &lock)
	}
}

var ID uint32 = 8

func startRead(token uint32, connMap map[uint32]*net.TCPConn, handerMap map[uint32]*ServerHandlerTrigger, lock *sync.Mutex) {
	data := make([]byte, 512)
	conn := connMap[token]
	handler := handerMap[token]
	ss := *handler
	for {
		size, err := conn.Read(data)
		if err != nil {
			utils.ZapLogger.Error(err.Error())
			return
		}
		if size == 0 {
			ss.OnClose()
			lock.Lock()
			delete(connMap, token)
			delete(handerMap, token)
			lock.Unlock()
			return
		}
		ss.OnMessage(data)
	}
}
