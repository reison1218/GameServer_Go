package tcp

import (
	"fmt"
	"net"
	"sync"
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
	clone() ServerHandlerTrigger
	onMessage(bytes []byte)
	onClose()
	onOpen(tcpConn TcpSession)
	getToken() uint32
}

type ServerHandler struct {
	Token uint32
	Conn  TcpSession
}

func (serverHandler ServerHandler) clone() ServerHandlerTrigger {
	return ServerHandler{serverHandler.Token, serverHandler.Conn}
}

func (serverHandler ServerHandler) onMessage(bytes []byte) {
	println("there is mess from client")
	str := string("hello client")
	ss := []byte(str)
	var _, err = serverHandler.Conn.Conn.Write(ss)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func (serverHandler ServerHandler) onClose() {
	println("client disconnect!")
}

func (serverHandler ServerHandler) onOpen(tcpConn TcpSession) {
	println("there is new client connect!")
	serverHandler.Token = tcpConn.Token
	serverHandler.Conn = tcpConn
}

func (serverHandler ServerHandler) getToken() uint32 {
	return serverHandler.Token
}

type TcpSession struct {
	Conn  *net.TCPConn
	Token uint32
}

func InitTcpServer(address string, serverHandler ServerHandlerTrigger) {
	tcpAddr, err := net.ResolveTCPAddr("tcp", address) //创建 tcpAddr数据
	if err != nil {
		fmt.Println(err)
		return
	}
	tcpListener, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		fmt.Println(err)
		return
	}

	tcpServer := newTcpServer(tcpListener, make(map[uint32]*net.TCPConn))

	var lock sync.Mutex
	fmt.Println("start Listen", tcpServer.tcpLister.Addr())
	handlerMap := make(map[uint32]*ServerHandlerTrigger)
	for {
		tcpConn, err := tcpServer.tcpLister.AcceptTCP()
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Println("there is new tcpClient coming", tcpConn.RemoteAddr())
		ID += 1
		ts := TcpSession{tcpConn, ID}
		sh := serverHandler.clone()
		sh.onOpen(ts)
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
			fmt.Println(err)
			return
		}
		if size == 0 {

			ss.onClose()
			lock.Lock()
			delete(connMap, token)
			delete(handerMap, token)
			lock.Unlock()
			return
		}
		ss.onMessage(data)
		fmt.Println("from client message:", data, "size:", size)
	}
}
