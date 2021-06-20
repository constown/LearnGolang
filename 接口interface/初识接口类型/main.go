package main

import "fmt"

// 接口是一种类型 是一种特殊的类型 规定了变量有哪些方法
// 引出接口的实例
// 在编程中会遇到以下场景:
// 我不关系一个变量是什么类型,只关心能调用他的什么方法

type speaker interface {
	speak()
}

type person struct {
}

type cat struct {
}

type dog struct {
}

func (c cat) speak() {
	fmt.Println("喵喵喵~")
}

func (d dog) speak() {
	fmt.Println("旺旺旺~")
}

func (p person) speak() {
	fmt.Println("嘤嘤嘤~")
}
func beat(x speaker) {
	// 接受一个参数,传进来谁就打谁
	x.speak()
}

func main() {
	var c1 cat
	var d1 dog
	var p1 person
	beat(c1)
	beat(d1)
	beat(p1)
}
