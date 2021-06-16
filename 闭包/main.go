package main

// 闭包
// 闭包是一个函数，这个函数包含了他外部作用的一个变量

// 底层的原理
// 1、函数可以作为返回值
// 2、函数内部查找变量的顺序，先在自己内部找，找不到往外层找
import (
	"fmt"
	"strings"
)

func f1(f func()) {
	fmt.Println("this is f1")
	f()
}
func f2(x, y int) {
	fmt.Println("this is f2")
	fmt.Println("f2调用结果：", x+y)
}

// 定义一个函数对f2进行包装
//func f3(f func(int, int)) func() {
//	tmp := func() {
//		f()
//	}
//	return tmp
//}
func main() {
	ret := add(f2, 100, 200)
	f1(ret)
	ret1 := adder()
	ret2 := ret1(200)
	fmt.Println("ret2:", ret2) // 300
	jpgFunc := makeSuffixFunc(".jpg")
	txtFunc := makeSuffixFunc(".txt")
	fmt.Println(jpgFunc("lucy"))
	fmt.Println(txtFunc("monkey.txt"))
}

func add(x func(int, int), m, n int) func() {
	tmp := func() {
		x(m, n)
	}
	return tmp
}

func adder() func(int) int {
	var x1 = 100
	return func(y int) int { // 这个返回的不仅是一个函数，还有 x1 变量的引用
		x1 += y
		return x1
	}
}

func makeSuffixFunc(suffix string) func(string) string {
	return func(name string) string {
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return name
	}
}
