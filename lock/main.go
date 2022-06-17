package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	sum()
}

// sum
func sum() {
	var num = 0
	var sync sync.Mutex
	for i := 0; i < 100; i++ {
		go func(i int) {
			sync.Lock()
			num++
			fmt.Printf("goroutine %d: num = %d\n", i, num)
			sync.Unlock()
		}(i)
	}
	time.Sleep(time.Second) //主goroutine等待1秒以确保所有工作goroutine执行完毕
}
