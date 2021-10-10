package main

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
)

//设置log参数
func init() {

	fmt.Println("main.init() run!")

	//设置输出样式，自带的只有两种样式logrus.JSONFormatter{}和logrus.TextFormatter{}
	logrus.SetFormatter(&logrus.JSONFormatter{})
	//设置output,默认为stderr,可以为任何io.Writer，比如文件*os.File
	logrus.SetOutput(os.Stdout)
	//设置最低loglevel
	logrus.SetLevel(logrus.DebugLevel)
}

func main() {

	fmt.Println("main.main() run!")

	//共有6个等级
	logrus.Debug("Useful debugging information.")
	logrus.Info("Something noteworthy happened!")
	logrus.Warn("You should probably take a look at this.")
	logrus.Error("Something failed but I'm not quitting.")
	logrus.Fatal("Bye.")   //log之后会调用os.Exit(1)
	logrus.Panic("I'm bailing.")   //log之后会panic()

	//使用field示例
	logrus.WithFields()
}
