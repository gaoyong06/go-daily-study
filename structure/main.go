package main

import "fmt"

// NewInt 类型定义
type NewInt int

// Myint 类型别名
type Myint = int

func main() {


	// 类型定义和类型别名的区别
	var a NewInt
	var b Myint

	fmt.Printf("type of a: %T\n", a) //type of a: main.NewInt
	fmt.Printf("type of b: %T\n", b) //type of b: int

	// Person 定义一个 人 类型的结构体
	// 方法一
	//type Person struct {
	//	Name string
	//	City string
	//	Age  int8
	//}

	//同类型的字段可以放在同一行

	// 方法二
	type Person struct {
		Name, City string
		Age  int8
	}

	//基本实例化

	//方式一
	//申明 Person 结构体类型变量 p1
	var p1 Person
	p1.Name = "张三"
	p1.City = "西安"
	p1.Age = 20

	fmt.Printf("p1=%v\n", p1)
	fmt.Printf("p1=%#v\n", p1)
	fmt.Println()

	//方式二
	p2 := Person{"张三1", "西安1", 21}
	fmt.Printf("p2=%v\n", p2)
	fmt.Printf("p2=%#v\n", p2)
	fmt.Println()

	//方式三
	p3 := Person{
		Name: "张三2",
		City: "西安2",
		Age: 22,
	}
	fmt.Printf("p3=%v\n", p3)
	fmt.Printf("p3=%#v\n", p3)
	fmt.Println()

	// 忽略的字段默认为 (零值)[../variable/note.md]
	// 下面是示例 Age:0
	p4 := Person{
		Name: "张三3",
		City: "西安3",
	}
	fmt.Printf("p4=%v\n", p4)
	fmt.Printf("p4=%#v\n", p4)
	fmt.Println()

}
