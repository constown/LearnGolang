package main

import (
	"fmt"
	"sync"
)

// channel 练习
// 1.启动一个goroutine，生成100个数发送到ch1
// 2.启动一个goroutine，从ch1中取出值，计算平方放到ch2
// 3.在main中 从ch2中取值打印
// 单向通道 chan<- 只能存值   <-chan 只能取值

var wg sync.WaitGroup
var once sync.Once

func f1(ch1 chan<- int) {
	defer wg.Done()
	for i := 0; i < 100; i++ {
		ch1 <- i
	}
	close(ch1) // 通道关闭了，不能存，但是可以取
}

func f2(ch1 <-chan int, ch2 chan<- int) {
	defer wg.Done()
	for {
		x, ok := <-ch1
		if !ok {
			break
		}
		ch2 <- x * x
	}
	once.Do(func() { close(ch2) }) // 确保操作只执行一次
}

func main() {
	wg.Add(3)
	a := make(chan int, 100)
	b := make(chan int, 100)
	go f1(a)
	go f2(a, b)
	go f2(a, b)
	for ret := range b {
		fmt.Println(ret)
	}
	wg.Wait()
}
