package main

import "fmt"

// 变量的全局作用域

// 定义一个全局作用域
var x = 100

// 定义一个函数
func f1() {
	x := 10
	// 变量的查找顺序
	// 1、先在函数内部查找
	// 2、找不到就往函数的外部查找，一直找到全局
	// 3、如果全局找不到，就报错你
	name := "Lucy"
	// 函数内部定义的变量无法在函数外部使用
	fmt.Println(x)
	fmt.Println(name)
}
func main() {
	f1() // 100
	//fmt.Println(name)

	// 语句块作用域
	// 只能在 if 语句中执行 for 同理
	if i := 10; i < 18 {
		fmt.Println("<18")
	}
	//fmt.Println(i)
}
