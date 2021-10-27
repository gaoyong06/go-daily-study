###### pointer 指针

指针是一种变量类型，他的特征是指针变量的值是内存地址。
区别于其他类型的变量，如数字类型，字符串类型的变量，数字类型变量的值是数字， 字符串类型的变量的值是字符串。

而变量是一个占位符，用于引用内存地址。
类似于变量名其实内存地址的“别名”(reference)

所以指针就是值是其他变量的内存地址的变量


- 使用场景

1. 传递比较大的数据结构时，可以选择中指针，避免值拷贝

多大的数据结构算比较大，并无定论，可以通过基准测试工具来做具体分析 //TODO 补充基准测试工具使用方法

2. 需要修改函数参数时需要传递该参数的指针

go语言中，函数(方法)是参数是值传递(值拷贝), 也就是拷贝了一个新的变量，作用域
在该函数体的。如果要在函数内部的修改外部变量的值，可以把外部的变量的地址传递进来。

```go
package main

import "fmt"

func fn1 (x int) {
	x = 10
}

func fn2(x *int)  {

	*x = 40
}


func main() {

	// 需要修改函数参数时需要传递该参数的指针
	x := 5
	fn1(x)
	fmt.Println("fn1 run end x = ", x) //fn1 run end x =  5
	fn2(&x)
	fmt.Println("fn2 run end x = ", x) //fn2 run end x =  40
}
```

3. api一致性

在定义函数（方法）时，如果传递是指针，则其他方法尽量也用指针，这将使得 API 更简单，避免去记到底哪里需要引用。

```go
type Persion struct {
	name string
	age int
	city string
}

func (p *Persion)rename(newName string)  {
	p.name = newName
}


// 推荐的方式
// 虽然为了一致性并不需要在 printName 中使用指针。但是这将使得 API 更简单，避免去记到底哪里需要引用。
func (p *Persion)printName()  {
	fmt.Println(p.name)
}

// 不推荐的方式
func (p Persion)anotherPrintName()  {

	fmt.Println(p.name)
}
```

4. 表示缺失

指针是引用类型，空值是nil, nil表示不存在。 非引用类型的空值是 0,false, "" 等, 非nil的零值,容易引起混淆。
例如学习的成绩是0, 0是表示该学生未参加考试,还是考试成绩是0

###### 指针变量

```go
// 申明实际变量
var a int = 1
// 申明指针变量
var b *int
```
一个指针变量通常缩写为 ptr


###### & 取地址

使用 (&) 取变量的地址

###### * 取值

在指针变量前面使用前缀 (*),获取该指针变量指向地址的值

```go
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
```

###### 空指针
当一个指针被定义后,没有被赋值，他的值是nil。
nil指针也称为空指针。


###### make

- 使用场景

###### new

- 使用场景


###### make 和 new的区别
