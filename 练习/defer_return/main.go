package main

import "fmt"

// 	go语言中 return 不是原子操作，在底层是分为两步执行
//	第一步：返回值赋值
//	第二步: 真正的ret返回
// 	函数中如果存在defer,那么在defer执行的时机是在第一步和第二步之间

func f1() int {
	x := 5
	defer func() {
		x++
	}()
	return x
}

func f2() (x int) {
	defer func() {
		x++
	}()
	return x
}

func f3() (y int) {
	x := 5
	defer func() {
		x++
	}()
	return x
}

func f4() (x int) {
	defer func(x int) {
		x++
	}(x)
	return 5
}

func main() {
	fmt.Println(f1()) // 5
	fmt.Println(f2()) // 6
	fmt.Println(f3()) // 5
	fmt.Println(f4()) // 5
}
