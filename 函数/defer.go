package main

import "fmt"

//  defer 把他后面的语句延迟到函数即将返回的时候执行
//  可以在释放一些资源的时候使用
//  一个函数中有多个defer语句
//	多个defer语句按照先进后出的顺序延迟执行

func deferDemo() {
	fmt.Println("start")
	defer fmt.Println("1")
	defer fmt.Println("2")
	defer fmt.Println("3")
	fmt.Println("end")
	//	start
	//	end
	//	3
	//	2
	//	1

}

func main() {
	deferDemo()
}
