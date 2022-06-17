package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func hello(i int) {
	defer wg.Done() // goroutine结束就登记-1
	fmt.Println("Hello Goroutine!", i)
}

func main() {

	//------
	// for i := 0; i < 10; i++ {
	// 	wg.Add(1) // 启动一个goroutine就登记+1
	// 	go hello(i)
	// }
	// wg.Wait() // 等待所有登记的goroutine都结束

	//-------
	// 合起来写
	go func() {
		i := 0
		for {
			fmt.Printf("new goroutine for run\n")
			i++
			fmt.Printf("new goroutine: i = %d\n", i)
			time.Sleep(time.Second)
		}
	}()

	i := 0
	for {
		i++
		fmt.Printf("main goroutine: i = %d\n", i)
		time.Sleep(time.Second)
		if i == 2 {
			break
		}
	}
}
