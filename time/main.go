package main

import (
	"fmt"
	"time"
)

func main() {

	time := time.Now().Add(time.Minute * -10).Unix()
	fmt.Println(time)

}
