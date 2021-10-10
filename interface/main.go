package main

import "fmt"

//定义一个猫的结构体
type cat struct {}

//猫实现Animal接口中的eat方法
func (c cat) eat() {
	fmt.Println("吃猫粮")
}

//定义一个狗的结构体
type dog struct {}

//狗实现Animal接口中的eat方法
func (d dog) eat() {
	fmt.Println("吃狗粮")
}

//定义一个人的结构体
type person struct {}

//人实现Animal接口中的eat方法
func (p person) eat()  {
	fmt.Println("吃饭")
}

// Animal 接口定义
type Animal interface {

	//接口中定义一个吃东西的方法
	eat()
}

//main里面定义一个吃东西的方法，接受Animal类型的参数
func eat(animal Animal)  {
	animal.eat()
}

func main() {
	c1 := cat{}
	//c1.eat()
	d1 := dog{}
	//d1.eat()
	p1 := person{}
	//p1.eat()
	eat(c1)
	eat(d1)
	eat(p1)

	var a1 Animal
	c2 := cat{}
	a1 = c2
	fmt.Println(a1)

}
