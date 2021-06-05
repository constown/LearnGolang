package main

import (
	"fmt"
	"math"
)

func main() {
	i1 := math.MaxFloat32 // float32 最大值
	fmt.Println(i1)

	f1 := 1.23456
	fmt.Printf("%T\n", f1) // 默认 go 语言中的小数都是 float64 类型

	f2 := float32(1.23456) // 显示声明为 float32 类型
	fmt.Printf("%T\n", f2)

	// f1 = f2  float32 类型的值不能直接赋值给 float64 类型的变量

	// go 语言中 bool 类型只有 true 和 false
	//	布尔类型变量默认值为 false
	// go 语言中不允许将整型强制转换为布尔型
	// 布尔型无法参与数值运算，也无法与其他类型进行转换

	b1 := true

	var b2 bool

	fmt.Printf("%T\n", b1)
	fmt.Printf("%T value:%v\n", b2, b2) // v 表示打印那个的值
}
