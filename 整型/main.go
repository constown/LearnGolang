package main

import "fmt"

// 在使用 int 和 uint 类型时，不能假定他是32位或者64位的整型，而是要考虑不同平台的差异
// uintptr 用于存放一个指针

func main() {
	//	go语言中无法直接定义二进制数

	// 十进制
	var i1 = 101
	fmt.Printf("%d\n", i1) // 十进制
	fmt.Printf("%b\n", i1) // 二进制
	fmt.Printf("%o\n", i1) // 八进制
	fmt.Printf("%x\n", i1) // 十六机制

	// 八进制
	i2 := 077
	fmt.Printf("%d\n", i2)

	// 十六进制
	var i3 = 0x1234567
	fmt.Printf("%d\n", i3)

	// 查看变量的类型
	fmt.Printf("%T\n", i3)

	//	声明 int8 类型的变量

	i4 := int8(9) // 明确指定int8类型，否则就是默认为int类型
	fmt.Printf("%T\n", i4)

}
