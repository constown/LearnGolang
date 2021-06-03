package main

import "fmt"

func main() {
	// 十进制
	var i1 = 101
	fmt.Printf("%d\n", i1) // 十进制
	fmt.Printf("%b\n", i1) // 二进制
	fmt.Printf("%o\n", i1) // 八进制
	fmt.Printf("%x\n", i1) // 十六机制

	// 八进制
	var i2 = 007
	fmt.Printf("%d\n", i2)

	// 十六进制
	var i3 = 0x1234567

	fmt.Printf("%x\n", i3)
}
