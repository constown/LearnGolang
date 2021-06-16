package main

import "fmt"

// 递归： 函数自己调用自己
// 递归适合处理那种 问题相同、问题规模 越来越小的场景
// 递归一定要有一个明确的退出条件

// 阶乘
// !3 = 3 * 2 * 1
// !4 = 4 * !3
// !5 = 5 * !4
func factorial(n uint64) uint64 {
	if n <= 1 {
		return 1
	}
	return n * factorial(n-1)
}

//有n台阶，一次可以走1步，也可以一次走2步, 一共有多少种走法
func onSteps(n uint64) uint64 {
	if n == 1 {
		return 1
	} else if n == 2 {
		return 2
	}
	return onSteps(n-1) + onSteps(n-2)
}

func main() {
	ret := factorial(5)
	fmt.Println("ret:", ret)
	res := onSteps(10)
	fmt.Println("res:", res)
}
