package main

import (
	"fmt"
	"time"
)

func main() {

	//10分钟前时间戳
	time1 := time.Now().Add(time.Minute * -10).Unix()
	fmt.Println(time1)

	//1天前时间戳
	theDayBefore := time.Now().AddDate(0, 0, -1).Unix()
	fmt.Println(theDayBefore)

}
