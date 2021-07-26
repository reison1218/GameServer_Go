package main

import (
	"fmt"
	"os"
	"sync"
	"time"
)

var (
	vname1 string
	vname2 string
)

func test(data *UserData) {
	data.userId += 100
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
func testDelete(t map[uint32]string) {
	delete(t, 1)
	fmt.Println("d", t)
}

const NUM_PRIMES uint32 = 50000

func sieve(v *[NUM_PRIMES]uint32, n *uint32, s *uint32) {
	var i uint32 = 0
	for i < *s {
		if *n%v[i] == 0 {
			*n++
			sieve(v, n, s)
			return
		}
		i++
	}
	v[*s] = *n
	*s++
}

func calc_n(n int) {
	print("N=", n)
	var ans int = 0
	var r1, r2, r3, r4, r5, r6 int
	var start = time.Now()
	for a1 := range make([]int, (n>>3)+1) {
		r1 = n - a1
		for a2 := range make([]int, (r1/7)+1) {
			r2 = r1 - a2
			for a3 := range make([]int, (r2/6)+1) {
				r3 = r2 - a3
				for a4 := range make([]int, (r3/5)+1) {
					r4 = r3 - a4
					for a5 := range make([]int, (r4>>2)+1) {
						r5 = r4 - a5
						for a6 := range make([]int, (r5/3)+1) {
							r6 = r5 - a6
							for a7 := range make([]int, (r6>>1)+1) {
								ans += a1 ^ a2 ^ a3 ^ a4 ^ a5 ^ a6 ^ a7 ^ (r6 - a7)
							}
						}
					}
				}
			}
		}
	}
	var end = time.Now()
	println(",", end.Sub(start).Milliseconds(), "ms")
}

func main() {
	calc_n(100)
	calc_n(160)
	calc_n(300)
	calc_n(400)
	calc_n(500)
	calc_n(600)
	//var size uint32 = 0
	//var i uint32 = 2
	//var primes [NUM_PRIMES]uint32
	//var start = time.Now()
	//for size < NUM_PRIMES {
	//	sieve(&primes, &i, &size)
	//}
	////for k := NUM_PRIMES-10; k < NUM_PRIMES; k++ {
	////	fmt.Printf("%d\n", primes[k])
	////}
	//var end = time.Now()
	//println(end.Sub(start).Milliseconds())
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
	//sh:=tcp.ServerHandler{0,tcp.TcpSession{nil,0}}
	//go tcp.InitTcpServer("127.0.0.1:8080",sh)
	//tcp.InitTcpClient("127.0.0.1:8080")
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

func testMutex(wait *sync.WaitGroup) {
	var lock sync.Mutex
	var j = new(Test)
	for i := 0; i < 999999; i++ {
		go add(j, &lock, wait)
	}
}

type Test struct {
	i int32
}

func add(j *Test, lock *sync.Mutex, wait *sync.WaitGroup) {
	lock.Lock()
	var i int32 = 0
	for ; i < 100; i++ {
		j.i += i
	}
	lock.Unlock()
	wait.Done()
}

func testSend(c chan int) {
	var startRecTime = time.Now()
	for i := 0; i < 999999; i++ {
		c <- 1
	}
	var endRecTime = time.Now()
	fmt.Println("send take time:", endRecTime.Sub(startRecTime))
}

func testRec(c chan int) {
	var startRecTime = time.Now()
	var res int = 0
	for i := range c {
		res += i
		if res >= 999999 {
			break
		}
	}
	var endRecTime = time.Now()
	fmt.Println("rec take time:", endRecTime.Sub(startRecTime), " i:", res)
}
