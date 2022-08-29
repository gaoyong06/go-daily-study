package main

import "fmt"

//因式分解的写法一般用于声明全局变量
var (
	a string
	b int
	c bool
)

func main() {

	//变量声明
	var a string
	fmt.Printf("a type: %T\n", a)
	fmt.Println()

	//变量初始化
	var d1 string = "hello world"
	fmt.Printf("d1 type: %T, d1 value: %v", d1, d1)
	fmt.Println()

	//零值
	var e1 bool   //布尔类型 false
	var e2 int    //数字类型 0
	var e3 string //字符串类型 ""

	var e4 *string          //指针类型 (*string)(nil)
	var e5 []int            //切片类型 (slice) []int(nil)
	var e6 map[string]int   //Map类型 map[string]int(nil)
	var e7 chan int         //通道类型 (channel) (chan int)(nil)
	var e8 func(string) int //函数类型 (func(string) int)(nil)
	var e9 error            //接口类型 (interface) nil

	var e10 struct{} //结构体类型 (struct) {}
	var e11 [5]int   //数组类型 [0,0,0,0,0]

	fmt.Printf("e1 type:%T, e1 value: %#v\n", e1, e1)
	fmt.Printf("e2 type:%T, e2 value: %#v\n", e2, e2)
	fmt.Printf("e3 type:%T, e3 value: %#v\n", e3, e3)
	fmt.Println()
	fmt.Printf("e4 type:%T, e4 value: %#v\n", e4, e4)
	fmt.Printf("e5 type:%T, e5 value: %#v\n", e5, e5)
	fmt.Printf("e6 type:%T, e6 value: %#v\n", e6, e6)
	fmt.Printf("e7 type:%T, e7 value: %#v\n", e7, e7)
	fmt.Printf("e8 type:%T, e8 value: %#v\n", e8, e8)
	fmt.Printf("e9 type:%T, e9 value: %#v\n", e9, e9)
	fmt.Println()
	fmt.Printf("e10 type:%T, e10 value: %#v\n", e10, e10)
	fmt.Printf("e11 type:%T, e11 value: %#v\n", e11, e11)
	fmt.Println()

	//变量类型推断
	var f1 = "hello world"
	fmt.Printf("f1 type:%T, f1 value: %#v\n", f1, f1)
	fmt.Println()

	//短变量声明
	//短变量声明只能用在函数内部
	//g1 := "hello world"
	//fmt.Printf("g1 type:%T, g1 value: %#v\n",g1, g1)

	//短变量声明和下面的代码是相同的

	//方式一
	var g1 string
	g1 = "hello world"
	fmt.Printf("g1 type:%T, g1 value: %#v\n", g1, g1)

	//方式二
	var g2 string = "hello world"
	fmt.Printf("g1 type:%T, g1 value: %#v\n", g2, g2)

	//方式三
	var g3 = "hello world"
	fmt.Printf("g1 type:%T, g1 value: %#v\n", g3, g3)

	//声明多个变量
	//方式一
	var h1, h2, h3 int
	h1, h2, h3 = 1, 2, 3
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
	fmt.Println()

	const (
		AvmType0 = iota
		AvmType1
		AvmType2
		AvmType3
	)

	fmt.Printf("AvmType0:%v, AvmType1:%v, AvmType2:%v, AvmType3:%v\n", AvmType0, AvmType1, AvmType2, AvmType3)

	if 3 < AvmType0 || 3 >= AvmType3 {

		fmt.Println("aaaaa")
	}

}
