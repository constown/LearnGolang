package main

import "fmt"

func f1(f func()) {
	fmt.Println("this is f1")
	f()
}
func f2(x, y int) {
	fmt.Println("this is f2")
	fmt.Println(x + y)
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
	ret()
}

func add(x func(int, int), m, n int) func() {
	tmp := func() {
		x(m, n)
	}
	return tmp
}
