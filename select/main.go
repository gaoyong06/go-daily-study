package main

import (
	"fmt"
	"time"
)

func test1(ch chan string) {
	time.Sleep(time.Second * 5)
	ch <- "test1"
}
func test2(ch chan string) {
	time.Sleep(time.Second * 2)
	ch <- "test2"
}

func recv1(output1 <-chan string, output2 <-chan string) {

	for {
		// 尝试从ch1接收值
		data1 := <-output1
		fmt.Println("data1=", data1)
		// 尝试从ch2接收值
		data2 := <-output2
		fmt.Println("data2=", data2)
		break
	}

}

func recv2(output1 <-chan string, output2 <-chan string) {
	// 用select监控
	select {
	case s1 := <-output1:
		fmt.Println("s1=", s1)
	case s2 := <-output2:
		fmt.Println("s2=", s2)
		// default:
		// 	// 如果上面都没有成功，则进入default处理流程
		// 	fmt.Println("default run")
	}
}

//-------------

// sameTimeEnd 两个通道同时结束
func sameTimeEnd() {
	// 创建2个管道
	int_chan := make(chan int, 1)
	string_chan := make(chan string, 1)
	go func() {

		int_chan <- 1
	}()
	go func() {
		// time.Sleep(2 * time.Second)
		string_chan <- "hello"
	}()
	select {
	case value := <-int_chan:
		fmt.Println("int:", value)
	case value := <-string_chan:
		fmt.Println("string:", value)
	}
}

//----------------

// isFull 判断管道有没有存满
func isFull() {

	// 创建管道
	output1 := make(chan string, 10)
	// 子协程写数据
	go write(output1)
	// 取数据
	for s := range output1 {
		fmt.Println("res:", s)
		time.Sleep(time.Second)
	}
}

func write(ch chan string) {
	for {
		select {
		// 写数据
		case ch <- "hello":
			fmt.Println("write hello")
		default:
			fmt.Println("channel full")
		}
		time.Sleep(time.Millisecond * 500)
	}
}

// func main() {
// 	// 2个管道
// 	// output1 := make(chan string)
// 	// output2 := make(chan string)
// 	// // 跑2个子协程，写数据
// 	// go test1(output1)
// 	// go test2(output2)

// 	// // recv1(output1, output2)
// 	// recv2(output1, output2)

// 	//---------
// 	// sameTimeEnd()

// 	//----------
// 	// isFull()

// 	fmt.Println("main结束")

// }

func main() {
	//closeChannel()
	c := make(chan int)
	timeout := time.After(time.Second * 2) //
	t1 := time.NewTimer(time.Second * 3)   // 效果相同 只执行一次
	var i int
	go func() {
		for {
			select {
			case <-c:
				fmt.Println("channel sign")
				return
			case <-t1.C: // 代码段2
				fmt.Println("3s定时任务")
			case <-timeout: // 代码段1
				i++
				fmt.Println(i, "2s定时输出")
			case <-time.After(time.Second * 4): // 代码段3
				fmt.Println("4s timeout。。。。")
			default: // 代码段4
				fmt.Println("default")
				time.Sleep(time.Second * 1)
			}
		}
	}()
	time.Sleep(time.Second * 6)
	close(c)
	time.Sleep(time.Second * 2)
	fmt.Println("main退出")
}
