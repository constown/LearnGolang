package main

import "fmt"

// 指针接收者实现接口和值接收者实现接口的区别
// 值接收者实现接口 结构体类型和结构体指针类型都能存
// 指针接收者实现接口 只能存结构体指针类型
type animal interface {
	move()
	eat(something string)
}

type cat struct {
	name string
	feet int8
}

// 方法使用值接收者
/*
func (c cat) move() {
	fmt.Println("cat走猫步")
}

func (c cat) eat(food string) {
	fmt.Printf("cat吃%s", food)
}
*/

// 使用指针接收者实现了接口的所有方法
func (c *cat) move() {
	fmt.Println("cat走猫步")
}

func (c *cat) eat(food string) {
	fmt.Printf("cat吃%s", food)
}

func main() {
	var a1 animal
	c1 := cat{"tom", 4}  // cat
	c2 := &cat{"假老练", 4} // 指针 *cat
	fmt.Println(c1)
	a1 = &c1
	fmt.Println(a1) // {tom 4}
	a1 = c2         // c1 c2都可以存进去
	fmt.Println(a1) // &{假老练 4}

}
