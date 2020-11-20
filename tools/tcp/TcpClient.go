package tcp

import (
	"fmt"
	"net"
	"time"
)

func InitTcpClient(address string){
	client, err := net.Dial("tcp", address)
	if err!=nil{
		fmt.Println(err)
		time.Sleep(time.Second*5)
		InitTcpClient(address)
	}
	startClientRead(client)
}

func startClientRead(clientConn net.Conn){
	data:=make([]byte,512)
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
