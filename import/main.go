package main

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"import/lib"
)

func main() {
	fmt.Println("hello world")
	lib.Example()
	logrus.Debug("debug")
}
