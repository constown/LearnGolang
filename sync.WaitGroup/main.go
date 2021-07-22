package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

/*
 * goroutine 结束的条件：
 * goroutine 对应的函数结束了，goroutine结束了
 * main函数执行完了，由main函数创建的goroutine也都结束了
 * 使用 sync.WaitGroup 来同步 goroutine 的执行
 */

func f() {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 5; i++ {
		r1 := rand.Int()    // int64
		r2 := rand.Intn(10) // 0 <= n < 10
		//fmt.Println(r1, r2)
		fmt.Printf("%T %v, %T %v\n", r1, r1, r2, r2)
	}
}

func f1(i int) {
	defer wg.Done()
	time.Sleep(time.Second * time.Duration(rand.Intn(3)))
	fmt.Println(i)
}

var wg sync.WaitGroup

func main() {
	// 启动10个goroutine,所有goroutine结束后才结束
	for i := 0; i < 5; i++ {
		wg.Add(1) // 每开启一个goroutine，计数1
		go f1(i)
	}
	wg.Wait() // 等待计数器结束
}
