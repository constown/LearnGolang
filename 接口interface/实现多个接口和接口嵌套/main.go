package main

import "fmt"

// 实现多个接口和接口嵌套

type animal interface {
	mover
	eater
}

type mover interface {
	move()
}

type eater interface {
	eat(string)
}

type cat struct {
	name string
	feet int8
}

// cat 实现了 mover 接口
func (c *cat) move() {
	fmt.Println("cat走猫步")
}

// cat 实现了 eater 接口
func (c *cat) eat(food string) {
	fmt.Printf("cat吃%s", food)
}

func main() {
	var a1 animal
	c1 := cat{
		name: "tom",
		feet: 4,
	}
	a1 = &c1
	fmt.Println(a1)
	a1.eat("fish")
}
