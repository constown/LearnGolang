package main

import (
	"fmt"
	"sync"
)

/*
Go 语言中的通道（channel）是一种特殊的类型。
通道像一个传送带或者队列，总是遵循先入先出（First In First Out）的规则，保证收发数据的顺序。
每一个通道都是一个具体类型的导管，也就是声明channel的时候需要为其指定元素类型。
*/

//var a []int
var b chan int // 需要指定通道中元素的类型， chan是一个引用类型，需要初始化才能使用
// make 初始化 slice，map， channel 通道必须初始化才能使用
var wg sync.WaitGroup

func noBufChannel() {
	fmt.Println(b)     // nil 没有初始化
	b = make(chan int) // 不带缓冲区的初始化通道
	wg.Add(1)
	go func() {
		defer wg.Done()
		x := <-b
		fmt.Println("后台goroutine从通道b中取到了x：", x)
	}()
	b <- 10
	fmt.Println("10发送到了通道b中")
	fmt.Println(b)
	wg.Wait()
	/*
		通道的操作： <-
		1、发送值  ch1 <- 1
		2、接收值  x := <- ch1
		3、关闭通道： close（)
	*/
}

func bufChannel() {
	fmt.Println(b)         // nil 没有初始化
	b = make(chan int, 16) // 不带缓冲区的初始化通道
	b <- 10
	x := <-b
	fmt.Println("10发送到了通道b中")
	fmt.Println(b, x)
}

func main() {
	//noBufChannel()
	bufChannel()
}
