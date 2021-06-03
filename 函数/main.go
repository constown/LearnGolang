package main

import (
	"fmt"
	"strings"
)

// 函数
//函数的定义
//func 函数名（参数）（返回值）{函数体}
//func sum(x int, y int) (ret int) {
//	return x + y
//}
//
//func main() {
//	r := sum(1, 2)
//	fmt.Println(r)
//}
// 词频统计
func main() {
	str := "how do you do"
	count := make(map[string]int)
	str2 := strings.Split(str, " ")
	for _, v := range str2 {
		count[v]++
	}
	fmt.Println(count)
}
