package main

import "fmt"

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
	Age        int8
}

// NewPerson 构造函数
func NewPerson(name, city string, age int8) *Person {

	return &Person{
		Name: name,
		City: city,
		Age:  age,
	}
}

func main() {

	// NewInt 类型定义
	type NewInt int

	// Myint 类型别名
	type Myint = int

	// 类型定义和类型别名的区别
	var a NewInt
	var b Myint

	fmt.Printf("type of a: %T\n", a) //type of a: main.NewInt
	fmt.Printf("type of b: %T\n", b) //type of b: int

	//结构体初始化

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

	//申明 Person 结构体类型变量 p3
	var p3 Person
	p3.Name = "张三"
	p3.City = "西安"
	p3.Age = 20

	fmt.Printf("p3=%v\n", p1)
	fmt.Printf("p3=%#v\n", p1)
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

	//匿名结构体
	var user struct {
		Name string
		Age  int
	}
	user.Name = "张三4"
	user.Age = 18
	fmt.Printf("%#v\n", user)

	//创建指针类型结构体
	var p5 = new(Person)
	fmt.Printf("%T\n", p5)
	fmt.Printf("p5=%#v\n", p5)

	//给结构体指针赋值
	p5.Name = "张三5"
	p5.Age = 19
	p5.City = "西安4"
	fmt.Printf("%T\n", p5)
	fmt.Printf("p5=%#v\n", p5)

	//取结构体的地址实例化
	//使用`&`对结构体进行取址操作相当于对该结构体类型进行了一次`new`实例化操作
	p6 := &Person{}
	fmt.Printf("%T\n", p6)
	fmt.Printf("p6=%#v\n", p6)
	p6.Name = "张三6"
	p6.Age = 20
	p6.City = "西安5"
	fmt.Printf("p6=%#v\n", p6)

	//面试题
	//下面的代码的执行结果是什么
	type student struct {
		name string
		age  int
	}

	m := make(map[string]*student)
	stus := []student{
		{name: "小王子", age: 18},
		{name: "娜扎", age: 23},
		{name: "大王八", age: 9000},
	}

	for _, stu := range stus {

		fmt.Println("======== Start =========")
		fmt.Println(stu)
		fmt.Printf("%#v\n", &stu)
		fmt.Println("======== End =========")
		m[stu.name] = &stu
	}

	// 为什么这里的value都是相同的?
	// map[大王八:0xc000004090 娜扎:0xc000004090 小王子:0xc000004090]
	// for _, stu := range stus, stu是一个临时变量，因为是值拷贝，所以每次循环都是将stus的一项值拷贝至stu,stu的内存地址是同一个，所以m[stu.name] = &stu
	// 后, m中的key虽然不同，但是值都是一个相同的内存地址，遍历完成后该内存地址存储的是stus的最后一项值, 所以在后面v.name中，打印的始终是相同的值
	//https://zhuanlan.zhihu.com/p/212828864
	//for _, v := range sliceX { use x} 问题在于v从始至终是一个v, 所以等遍历完之后v被赋值成了o最后一个数据, 你所有的oPointer里面存的是一个指向v的指针, 所以你打印出来都是v的内容就不奇怪了. for range 使用指针, 还有go func {} 捕获参数的时候都有这个陷阱. 写代码的时候保持对陷阱的警惕.
	fmt.Println(m)

	for k, v := range m {
		fmt.Println(k, "=>", v.name)
	}

	//构造函数
	//调用沟通函数
	p7 := NewPerson("张三7", "西安7", 27)
	fmt.Println(p7)

	//"爸爸的爸爸是爷爷"
	type People struct {
		name string

		////这里要用指针类型,否则编译器会报invalid recursive type People的错误
		//https://stackoverflow.com/questions/59935466/invalid-recursive-type-and-illegal-cycle-in-declaration-of
		child *People
	}

	relation := &People{
		name: "爷爷",
		child: &People{
			name: "爸爸",
			child: &People{
				name: "我",
			},
		},
	}
	fmt.Println(relation)

	//方法和接受者

}
