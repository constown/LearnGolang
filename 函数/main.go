package main

import (
	"fmt"
	"strings"
)

// 函数
//函数的定义
//func 函数名（参数）（返回值）{函数体}
func sum(x int, y int) (ret int) {
	return x + y
}

// 没有返回值
func f1(x int, y int) {
	fmt.Println(x + y)
}

// 没有参数没有返回值
func f2() {
	fmt.Println("f2")
}

// 没有参数有返回值
func f3() int {
	return 3
}

// 参数可以命名，也可以不命名
// 命名的返回值就相当于在函数中声明一个变量，在函数中就可以直接使用，不需要再声明
// 对比 f4 和 f5 的区别
func f4(x int, y int) (ret int) {
	ret = x + y
	return // 使用命名返回可以省略return后面的值
}
func f5(x int, y int) int {
	ret := x + y
	return ret
}

// 多个返回值
func f6() (int, string) {
	return 1, "沙河"
}

// 参数的类型简写, 当参数中连续多个参数的类型一致时候，我们可以省略非最后一个参数的类型
func f7(x, y, z int, m, n string) int {
	return x + y
}

// 可变长参数
// 可变参数必须放在函数参数的最后一个
func f8(x string, y ...int) {
	fmt.Println(x)
	fmt.Println(y) // 打印的是一个切片 []int 切片类型是根据参数定义来的
}

// go语言中函数没有默认参数这个概念

// 词频统计
func main() {
	str := "how do you do"
	count := make(map[string]int)
	str2 := strings.Split(str, " ")
	for _, v := range str2 {
		count[v]++
	}
	fmt.Println(count)

	r := sum(1, 2)
	fmt.Println(r)

	_, n := f6()
	fmt.Println(n)

	f8("下雨了", 1, 2, 3, 4, 5)
}
