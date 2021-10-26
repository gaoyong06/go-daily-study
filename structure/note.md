###### 示例程序
[main.go](./main.go)

###### 结构体 structure

Go语言没有“类”的概念，结构体有类似其他语言中“类”的特征，且通过结构体实现比其他语言更高的灵活性，Go语言中通过struct来实现面向对象

###### 类型
Go语言中有一些基本数据类型，如string,整形，浮点型，布尔，等数据类型。
GO语言中可以使用`type` 关键字来定义自定义类型。

###### 自定义类型

自定义类型是定义一个全新的类型。可以基于内置的基本类型定义，也可以通过struct, interface定义。

```go
//将MyInt定义为int类型
type MyInt int

//定义Person结构体
type Person struct {
    
    name string,
    age  int
}

//定义Animal[接口](../interface/note.md)
type Animal Interface {
    fun move()
}
```

###### 类型别名

类型别名是`Go1.9`版本添加的新功能
类型别名规定：TypeAlias只是Type的别名，本质上TypeAlias于Type是同一个类型。
```go
type TypeAlias = Type
```

###### 类型定义和类型别名的区别

从语义上讲，类型定义是定义一个新类型，类型别名是给类型的另一个称呼，类型别名本质上没有产生新类型。

```go
package main

import "fmt"

// NewInt 类型定义
type NewInt int

// Myint 类型别名
type Myint = int

func main() {
	
	var a NewInt
	var b Myint

	fmt.Printf("type of a: %T\n", a) //type of a: main.NewInt
	fmt.Printf("type of b: %T\n", b) //type of b: int
}
```
a 的类型是 `main.NewInt`, b 的类型是 `int`. `MyInt`类型只会在代码中存在，编译完成时并不会有`MyInt`类型。


###### 结构体
结构体是一种自定义数据类型

###### 结构体定义

```go
type 类型名 struct {
    
    字段名 字段类型
    字段名 字段类型
    ...
    字段名 字段类型
}
```

类型名：标识自定义结构体的名称，在同一个包内不能重复。
字段名：表示结构体字段名。结构体中的字段名必须唯一。
字段类型：表示结构体字段的具体类型。

定义一个`Persion`(人)类型的结构体

```go
// Person 定义一个 人 类型的结构体
// 方法一
type Person struct {
    Name string
    City string
    Age  int8
}

//同类型的字段可以放在同一行
// 方法二
type Person struct {
    Name, City string
    Age  int8
}
```

内置的基础数据类型是用来描述**一个**值的，而结构体是用来描述**一组**值的。


###### 结构体实例化

结构体的定义只是一种内存布局描述，
只有当结构体实例化时，才会真正地分配内存。也就是是必须实例化后才能使用结构体字段。

###### 普通实例化
结构体本身是一种类型，可以像整形、字符串等类型一样，以`var`的方式声明结构体即可完成实例化。
这种用法是为了更明确地表示一个变量被设置为零值。

```go
// p1为结构体实例,Person为结构体类型
var p1 Person
```

###### new 实例化
`go`语言中，可以使用`new`关键字对类型（包括结构体，整形，浮点数，字符串等）进行实例化，
结构体在实例化后会形成指针类型的结构体。

```go
var p5 = new(Person) 
fmt.Printf("p5=%#v\n", p5) //p5=&main.Person{Name:"", City:"", Age:0}
```

###### & 取地址实例化
在`Go`语言中，对结构体进行`&`取地址操作时，视为对该类型进行一次`new`的实例化操作

```go
p6 := &Person{}
fmt.Printf("p6=%#v\n", p6) //p6=&main.Person{Name:"", City:"", Age:0}
```

###### 结构体体零值

结构体的零值是结构体中的每个字段都是零值

```go
var p1 Person
fmt.Printf("p1=%#v\n", p1) //p1=main.Person{Name:"", City:"", Age:0}
fmt.Println()
```
###### 结构体初始化

结构体在实例化时可以直接对成员变量进行初始化，初始化有两种形式分别是以字段"**键值对**"形式和"**多个值**"得列表形式，
"**键值对**"形式的初始化适合选择性填充字段较多的结构体，"**多个值**"的列表形式适合填充字段较少的结构体。

```go
//键值对形式
//方式一
p1 := Person{
Name: "张三2",
City: "西安2",
Age:  22,
}
fmt.Printf("p1=%v\n", p1)
fmt.Printf("p1=%#v\n", p1)
fmt.Println()

//值列表形式
//方式二
p2 := Person{"张三1", "西安1", 21}
fmt.Printf("p2=%v\n", p2)
fmt.Printf("p2=%#v\n", p2)
fmt.Println()
```

值列表初始化，需注意：
- 必须初始化结构体的所有字段
- 每一个初始值得填充顺序必须与字段在结构体中的申明顺序一致
- 键值对与值列表的初始化形式不能混用

###### 匿名结构体
匿名结构体没有类型名称,无需通过`type`关键字定义就可以直接使用

###### 匿名结构体定义和初始化
```go
var user struct {
    Name string
    Age  int
}
user.Name = "张三4"
user.Age = 18
fmt.Printf("%#v\n", user)
```




###### 方法和接收者
Go语言中的`方法(Method)`是一种作用于特定类型变量的函数。 这种特定类型变量叫做`接受者（Receiver）`。接受者的概念就类似于其他语言中的`this`或者`self`

方法于函数的区别是，函数不属于任何类型，方法属于特定的类型。

###### 任意类型添加方法

**注意事项：** 非本地类型不能定义方法，也就是说我们不能给别的包的类型定义方法

###### 结构体的匿名字段

###### 嵌套结构体

###### 嵌套匿名字段

###### 嵌套结构体的字段名冲突


###### 实践
一个结构体，结构体字段类型 是一个接口
D:\work\service_core\service\game.go
line:61

###### 设计思想