package main

import "fmt"

// 结构体模拟实现继承

type animal struct {
	name string
}

// 给animal实现一个移动的方法
func (a animal) move() {
	fmt.Printf("%s会动\n", a.name)
}

type dog struct {
	feet uint8
	animal
}

// 给狗一个汪汪汪的方法
func (d dog) wang() {
	fmt.Printf("%s汪汪汪\n", d.name)
}
func main() {
	d1 := dog{
		feet: 4,
		animal: animal{
			name: "小明",
		},
	}
	fmt.Println(d1)
	d1.wang()
	d1.move()
}
