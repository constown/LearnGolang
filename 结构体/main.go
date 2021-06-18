package main

import "fmt"

// 结构体
// 可以理解为 js中的对象？

type person struct {
	name   string
	age    int
	gender string
	hobby  []string
}

func main() {
	var lucy person
	// 通过字段来赋值
	lucy.name = "lucy"
	lucy.age = 18
	lucy.gender = "女"
	lucy.hobby = []string{"足球", "篮球", "双色球"}
	fmt.Printf("类型：%T value：%v\n", lucy, lucy)

	// 匿名结构体 多用于临时场景
	var s struct {
		x string
		y int
	}

	s.x = "嘿嘿"
	s.y = 180
	fmt.Printf("类型：%T value：%v\n", s, s)
}
