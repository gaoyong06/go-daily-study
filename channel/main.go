package main

import (
	"fmt"
	"time"
)

// 使用无缓冲区(阻塞的)通道

// 接收消息
func recv(c chan int) {
	time.Sleep(3 * time.Second)
	ret := <-c
	fmt.Println("接收成功", ret)
}

// 阻塞发送
func sendBlock() {
	ch := make(chan int)
	go recv(ch) // 启用goroutine从通道接收值
	fmt.Println("开始发送")
	ch <- 10 //阻塞
	fmt.Println("发送成功")
	fmt.Println("执行结束")
}

// 非阻塞发送
func sendNonBlock() {

	ch := make(chan int, 1)
	go recv(ch) // 启用goroutine从通道接收值
	fmt.Println("开始发送")
	ch <- 10
	ch <- 11
	ch <- 12
	fmt.Println("发送成功")
	fmt.Println("执行结束")
}

// func main() {

// 	// sendBlock()
// 	sendNonBlock()
// }

// channel 练习
func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	// 开启goroutine将0~100的数发送到ch1中
	go func() {
		for i := 0; i < 100; i++ {

			fmt.Println("发送 ", i, "至ch1")
			ch1 <- i
		}
		close(ch1)
	}()
	// 开启goroutine从ch1中接收值，并将该值的平方发送到ch2中
	go func() {
		for {
			i, ok := <-ch1 // 通道关闭后再取值ok=false
			if !ok {
				break
			}
			fmt.Println("发送 ", i*i, "至ch2")
			ch2 <- i * i
		}
		close(ch2)
	}()
	// 在主goroutine中从ch2中接收值打印
	for i := range ch2 { // 通道关闭后会退出for range循环

		fmt.Println("接收 ch2 ", i)
	}
}
