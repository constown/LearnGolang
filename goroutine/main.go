package main

import (
	"fmt"
	"time"
)

// goroutine 并发编程
func hello(i int) {
	fmt.Println("hello", i)
}

/*
 * goroutine 结束的条件：
 * goroutine 对应的函数结束了，goroutine结束了
 * main函数执行完了，由main函数创建的goroutine也都结束了
 */

// 程序启动之后，会创建一个主goroutine去执行
func main() {
	for i := 0; i < 100000; i++ {
		go hello(i)
	}
	// 开启一个单独的goroutine去执行hellO函数
	fmt.Println("main")
	time.Sleep(time.Second)
	// main 函数结束了 由main函数启动的goroutine也都结束了
}
