package main

import "fmt"

// 全局变量
var (
	name string
)

// 在函数中声明的是局部变量

func main() {
	name = "小明"
	fmt.Print(name)
	// ↑是先声明后赋值

	// 声明变量的同时赋值
	var age int = 46
	fmt.Println(age)

	// 类型推导（根据值自动判断变量的类型）

	var isOK = true
	fmt.Print(isOK)

	// 简短变量声明, 只能在函数中使用， 不能在函数外声明
	s3 := "hhhh"
	fmt.Print(s3)
}
