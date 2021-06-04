package main

import "fmt"

//	常量是在程序运行中恒定不变的量，使用关键字 const 进行声明
//	常量一般定义在全局
const pi = 3.14115926

//	常量也可以批量声明
const (
	code  = 200
	state = 404
)

// 在批量声明的时候，如果某一行没有写值，就默认为和上一行相同
const (
	n1 = 100
	n2 // 100
	n3 // 100
)

// iota 是go语言的计数器 只能在常量的表达式中使用
// iota 在 const 关键字出现时 将被重置为 0 ， const 中每新增一行常量声明将 使 iota  的计数增加一次

const (
	a1 = iota // 0
	a2 = iota // 1
	a3 = iota // 2
)

// 插队，可以把 iota 看做行索引
const (
	b1 = iota // 0
	b2        // 1
	_         //2
	b3        //3
)

const (
	c1 = iota // 0
	c2 = 100  // 100
	c3 = iota // 2
	c4        // 3
)

// 多个常量声明在一行
const (
	d1, d2 = iota + 1, iota + 2 // 1 2
	d3, d4 = iota + 1, iota + 2 // 2 3
)

// 定义数量级
const (
	_  = iota
	KB = 1 << (10 * iota) // << 表示左移
	MB = 1 << (10 * iota)
	GB = 1 << (10 * iota)
	TB = 1 << (10 * iota)
	PB = 1 << (10 * iota)
)

func main() {
	//	pi = 23 常量定义之后就不能修改 修改就会报错
	fmt.Println("pi:", pi)
	fmt.Println("code:", code)
	fmt.Println("state：", state)

	fmt.Println("n1:", n1)
	fmt.Println("n2:", n2)
	fmt.Println("n3:", n3)

	fmt.Println("a1:", a1)
	fmt.Println("a2:", a2)
	fmt.Println("a3:", a3)

	fmt.Println("b1:", b1)
	fmt.Println("b2:", b2)
	fmt.Println("b3:", b3)

	fmt.Println("c1:", c1)
	fmt.Println("c2:", c2)
	fmt.Println("c3:", c3)
	fmt.Println("c4:", c4)

	fmt.Println("d1:", d1)
	fmt.Println("d2:", d2)
	fmt.Println("d3:", d3)
	fmt.Println("d4:", d4)

	fmt.Println("KB:", KB)
	fmt.Println("MB:", MB)
	fmt.Println("GB:", GB)
	fmt.Println("PB:", PB)
}
