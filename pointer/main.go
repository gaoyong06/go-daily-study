package main

import "fmt"

func fn1(x int) {
	x = 10
}

func fn2(x *int) {

	*x = 40
}

type Persion struct {
	name string
	age  int
	city string
}

func (p *Persion) rename(newName string) {
	p.name = newName
}

// 推荐的方式
// 虽然为了一致性并不需要在 printName 中使用指针。但是这将使得 API 更简单，避免去记到底哪里需要引用。
func (p *Persion) printName() {
	fmt.Println(p.name)
}

// 不推荐的方式
func (p Persion) anotherPrintName() {

	fmt.Println(p.name)
}

type exam struct {

	// 成绩
	score int
	// 是否参加考试
	// bool的零值是false
	// *bool的零值是nil
	present bool
}

func main() {

	// 申明实际变量
	var a int = 1
	// 申明指针变量
	var b *int

	// 实际变量的值
	fmt.Printf("a 变量的值是：%x\n", a)
	//为指针变量赋值
	b = &a
	fmt.Printf("a 变量的地址是：%x\n", &a)

	//指针变量的值
	fmt.Printf("b 指针变量的值是：%v\n", b)

	//访问指针变量中指向地址的值
	fmt.Printf("b 指针变量指向地址的值：%d\n", *b)

	// 需要修改函数参数时需要传递该参数的指针
	x := 5
	fn1(x)
	fmt.Println("fn1 run end x = ", x) //fn1 run end x =  5
	fn2(&x)
	fmt.Println("fn2 run end x = ", x) //fn2 run end x =  40

	//api一致性
	p1 := Persion{
		name: "张三",
		age:  18,
		city: "西安",
	}
	fmt.Println(p1)

	p1.rename("张三1")
	p1.printName()

	p1.rename("张三2")
	p1.anotherPrintName()

	//表示缺失
	e1 := exam{}
	fmt.Println(e1)

	//空指针
	var ptr *int
	fmt.Printf("ptr的值为：%v\n", ptr)
	fmt.Printf("ptr的地址为：%x\n", &ptr)

	//判断空指针
	if ptr == nil {
		fmt.Println("ptr == nil")
	}

	// 待研究
	// https://www.jianshu.com/p/97bfe8104e03
	var pi *int = nil
	var i interface{}
	i = pi
	fmt.Printf("pi的值为：%#v\n", pi)
	fmt.Printf("pi的地址为：%x\n", &pi)
	fmt.Printf("i的值为：%#v\n", i)
	fmt.Printf("i的地址为：%x\n", &i)
	fmt.Println(i == nil) // 结果为 false

}
