package main

import (
	"fmt"
	"net"
	"time"
)

type TcpServer struct {
	tcpLister *net.TCPListener
	tcpConnMap map[uint32]ServerHandler
}

type ServerHandlerTrigger interface {
	onMessage(bytes []byte)
	onClose()
	onOpen(tcpConn *net.TCPConn)
}

type ServerHandler struct {
	token uint32
	conn *net.TCPConn
}

func (serverHandler ServerHandler) onMessage(byte []byte)  {

}

func (serverHandler ServerHandler) onClose()  {

}

func (serverHandler ServerHandler) onOpen(tcpConn *net.TCPConn)  {

}

func initTcpServer(address string){
	tcpAddr, err := net.ResolveTCPAddr("tcp", address) //创建 tcpAddr数据
	tcplistener, err := net.ListenTCP("tcp", tcpAddr)
	if err !=nil{
		fmt.Println(err)
		return
	}
	tcpServer:=new(TcpServer)
	tcpServer.tcpLister = tcplistener
	tcpServer.tcpConnMap = make(map[uint32]ServerHandler)

	fmt.Println("start Listen",tcpServer.tcpLister.Addr())
	for{
		tcpConn,error:=tcpServer.tcpLister.AcceptTCP()
		if error !=nil{
			fmt.Println(error)
			continue
		}
		fmt.Println("there is new tcpClient coming",tcpConn.RemoteAddr())
		ID+=1
		serverHandler:=ServerHandler{ID,tcpConn}
		serverHandler.onOpen(tcpConn)
		tcpServer.tcpConnMap[ID] = serverHandler
		tcpConnRes := tcpServer.tcpConnMap[ID]
		go startRead(&tcpConnRes)
	}
}

var ID uint32 = 8

func startRead(map *map[uint32]ServerHandler,serverHandler *ServerHandler){
	data:=make([]byte,2014)
	var conn = serverHandler.conn
	for{
		size,error := conn.Read(data)
		if error !=nil{
			return
		}
		if size == 0{
			serverHandler.onClose()
			delete(map,serverHandler.token)
			return
		}
		serverHandler.onMessage(data)
		fmt.Println("from client message:",data,"size:",size)

		//str := string("hello client")
		//var b []byte =[]byte(str)
		//conn.Write(b)
	}
}

func initTcpClient(){
	//service := "127.0.0.1:8080"
	//tcpAddr, err := net.ResolveTCPAddr("tcp", service) //创建 tcpAddr数据
	//if err !=  nil{
	//	fmt.Println(err)
	//	return
	//}

	client, err := net.Dial("tcp", "localhost:8080")
  if err!=nil{
  	fmt.Println(err)
  	time.Sleep(time.Second*5)
  	initTcpClient()
  }
	startClientRead(client)
}

func startClientRead(clientConn net.Conn){
	data:=make([]byte,2014)
	for{
		str:=string("hello server,i am client")
		res:=[]byte(str)
		clientConn.Write(res)
		size,error := clientConn.Read(data)
		if error != nil{
			fmt.Println(error)
			continue
		}
		fmt.Println("from server mess:",data,"size",size)
	}
}