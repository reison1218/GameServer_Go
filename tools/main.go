package main

import (
	"fmt"
	"os"
	"sync"
	"time"

	"./tcp"
)

var (
	vname1 string
	vname2 string
)

func test(data *UserData){
	data.userId+=100
}

func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal error :", err.Error())
		os.Exit(1)
	}
}


func typeof(v interface{}) {
	fmt.Printf("type is: %T\n", v)

}
func testDelete(t map[uint32]string){
	delete(t,1)
	fmt.Println("d",t)
}
func main() {
	//var c = make(chan int)
	//go testSend(c)
	//testRec(c)
	//test := make(map[string]string)
	//test["1"] = "1"
	//var str string =test["1"]
	//str = "2"
	//println(str)
	//for k,v:= range test{
	//	println(k,v)
	//}
	sh:=tcp.ServerHandler{0,tcp.TcpSession{nil,0}}
	go tcp.InitTcpServer("127.0.0.1:8080",sh)
	tcp.InitTcpClient("127.0.0.1:8080")
	//a := time.Now().Nanosecond()
	//i := 1
	//for i < 10000000 {
	//	i += i
	//}
	//println(`最终值：`, i)
	//println(`时间：`, time.Now().Nanosecond()-a, `纳秒`)



	//var wait sync.WaitGroup
	//wait.Add(999999)
	//var start = time.Now()
	//testMutex(&wait)
	//wait.Wait()
	//var end = time.Now()
	//fmt.Println("耗时：",end.Sub(start))
}

func testMutex(wait *sync.WaitGroup){
	var lock  sync.Mutex
	var j  = new(Test)
	for i:=0;i<999999;i++{
		go add(j,&lock,wait)
	}
}

type Test struct {
	i int32
}
func add(j *Test,lock *sync.Mutex,wait *sync.WaitGroup){
	lock.Lock()
	var i int32 = 0
	for ;i<100;i++{
		j.i+=i
	}
	lock.Unlock()
	wait.Done()
}

func testSend(c chan int){
	var startRecTime = time.Now()
	for i:= 0;i< 999999;i++{
		c<-1
	}
	var endRecTime = time.Now()
	fmt.Println("send take time:",endRecTime.Sub(startRecTime))
}

func testRec(c chan int){
	var startRecTime = time.Now()
	var res int = 0
	for i := range c{
		res+=i
		if res >= 999999{
			break
		}
	}
	var endRecTime = time.Now()
	fmt.Println("rec take time:",endRecTime.Sub(startRecTime)," i:",res)
}


