package lib

import "fmt"

//init 在包导入时自动执行
//没有参数也没有返回值
//多用于一些初始化操作
func init()  {
	fmt.Println("")
}

func Example() {

	println(" lib.Example() is run! ")
}


