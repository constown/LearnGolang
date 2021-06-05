package main

import "fmt"

func main() {
	//查看类型
	n := 100
	fmt.Errorf("%T\n", n) // 查看类型
	fmt.Printf("%v\n", n) // 查看变量的值
	fmt.Printf("%d\n", n) // 十进制
	fmt.Printf("%b\n", n) // 二进制
	fmt.Printf("%o\n", n) // 八进制
	fmt.Printf("%x\n", n) // 十六进制

	s := "hello"
	fmt.Printf("字符串：%s\n", s)  // 字符串
	fmt.Printf("字符串：%v\n", n)  // 查看变量的值
	fmt.Printf("字符串：%#v\n", n) // 查看变量的值 （加个双引号）

}
