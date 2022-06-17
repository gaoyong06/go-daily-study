###### 示例程序
[main.go](./main.go)

###### variable 变量

变量是计算机语言中存储计算结果或表示值得抽象概念
类似于给一个内存地址起了个名字

Go语言变量名由字母、数字、下划线组成，其中首个字符不能为数字。

###### 变量声明

即声明一个变量

var 变量名 变量类型
示例：
```
var a string
```

一次声明多个变量

var 变量名1, 变量名2, 变量名3 变量类型
示例：
```
var a1, b1 int
```

###### 变量初始化

即声明一个变量并设置初始值

var 变量名 变量类型 = 值
示例：
```
var d1 string = "hello world"
```

###### 零值

在Go语言中如果声明了一个变量但是没有对它进行赋值操作，那么这个变量就会有一个类型的默认零值
nil是预定义的标识符，代表指针、通道、函数、接口、映射或切片的零值

- 数值类型零值为0
- 浮点类型零值为0.1
- 布尔类型零值为false
- 字符字符串类型零值为""
- 指针、函数、接口、切片、通道、映射类型的零值为nil

 bool:false 
 integer:0
 float:0.0 ,
 string:"" , 
 pointer, function, interface, slice, channel, map: nil 

```
	//零值
	var e1 bool					//布尔类型 false
	var e2 int					//数字类型 0
	var e3 string				//字符串类型 ""

	var e4 *string				//指针类型 (*string)(nil)
	var e5 []int				//切片类型 (slice) []int(nil)
	var e6 map[string]int 		//Map类型 map[string]int(nil)
	var e7 chan int				//通道类型 (channel) (chan int)(nil)
	var e8 func(string) int		//函数类型 (func(string) int)(nil)
	var e9 error				//接口类型 (interface) nil

	var e10 struct {}			//结构体类型 (struct) {}
	var e11 [5]int 				//数组类型 [0,0,0,0,0]

	fmt.Printf("e1 type:%T, e1 value: %#v\n",e1, e1)
	fmt.Printf("e2 type:%T, e2 value: %#v\n",e2, e2)
	fmt.Printf("e3 type:%T, e3 value: %#v\n",e3, e3)
	fmt.Println()
	fmt.Printf("e4 type:%T, e4 value: %#v\n",e4, e4)
	fmt.Printf("e5 type:%T, e5 value: %#v\n",e5, e5)
	fmt.Printf("e6 type:%T, e6 value: %#v\n",e6, e6)
	fmt.Printf("e7 type:%T, e7 value: %#v\n",e7, e7)
	fmt.Printf("e8 type:%T, e8 value: %#v\n",e8, e8)
	fmt.Printf("e9 type:%T, e9 value: %#v\n",e9, e9)
	fmt.Println()
	fmt.Printf("e10 type:%T, e10 value: %#v\n",e10, e10)
	fmt.Printf("e11 type:%T, e11 value: %#v\n",e11, e11)
```

###### 变量的类型推断
根据类型的值自动推断变量类型

var 变量名 = 值
示例：
```
var f1 = "hello world"
```

###### 短变量声明
函数体内声明局部变量的一种简写方式
示例：
```
g1 := "hello world"
```

###### 多变量声明
即一次声明多个变量

```
//因式分解的写法一般用于声明全局变量
var (
	a string
	b int
	c bool
)

func main() {

    //方式一
	var h1, h2, h3 int
	h1, h2, h3 = 1,2,3
	fmt.Printf("h1:%v, h2:%v, h3:%v\n", h1, h2, h3)

	//方式二
	var h4, h5, h6 = 4, 5, 6
	fmt.Printf("h4:%v, h5:%v, h6:%v\n", h4, h5, h6)

	//方式三
	h7, h8, h9 := 7, 8, 9
	fmt.Printf("h7:%v, h8:%v, h9:%v\n", h7, h8, h9)

	//方式四
	//上面5~10行处的全局代码声明
	a = "hello"
	b = 1
	c = true
	fmt.Printf("a:%v, b:%v, c:%v\n", a, b, c)
}
```


###### 容易出现的一些错误

```go
package main

import "fmt"

func main() {
   var a string = "abc"
   fmt.Println("hello, world")
}
```


//https://blog.csdn.net/weixin_42137723/article/details/112099869


###### 变量的生命周期
https://www.cnblogs.com/heych/p/12577492.html
https://www.jianshu.com/p/78f10bdbac73


