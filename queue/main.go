package main

import (
	"fmt"

	"github.com/gaoyong06/go-daily-study/queue/queue"
)

func main() {

	var queue queue.Queue

	//向队列末尾添加元素
	queue.Put(1)
	queue.Put(2)
	queue.Put(3)
	queue.Put(4)
	queue.Put(5)
	queue.Put(6)

	//判断队列是否为空
	fmt.Printf("queue.Empty: %#v\n", queue.Empty())

	//获取队列中的数据
	fmt.Printf("queue.GetItems: %#v\n", queue.GetItems())

	//获取队列长度
	fmt.Printf("queue.Len: %#v\n", queue.Len())

	//获取对首元素但不从队列中删除
	item, err := queue.Peek()
	fmt.Printf("queue.Peek. item: %#v, err:%#v\n", item, err)
	fmt.Printf("queue.Len: %#v\n", queue.Len())

	//阻塞获取对列前number元个素并从队列删除
	items, err := queue.Poll(2, 0)
	fmt.Printf("queue.Poll. items: %#v, err:%#v\n", items, err)
	fmt.Printf("queue.GetItems: %#v\n", queue.GetItems())
	fmt.Printf("queue.Len: %#v\n", queue.Len())

	//释放队列资源
	// items = queue.Dispose()
	// fmt.Printf("queue.Dispose. items: %#v\n", items)
	// fmt.Printf("queue.GetItems: %#v\n", queue.GetItems())
	// fmt.Printf("queue.Len: %#v\n", queue.Len())
	// items = queue.Dispose()
	// fmt.Printf("queue.Dispose. items: %#v\n", items)

	//获取队列中元素并删除直到checker返回false
	items, err = queue.TakeUntil(checker)
	fmt.Printf("queue.TakeUntil. items: %#v, err:%#v\n", items, err)

}

func checker(item interface{}) bool {

	fmt.Printf("checker item: %#v\n", item)
	return item == 3 || item == 4
}
