package main

import "fmt"

func main() {
	// 数组
	// 存放元素的容器，必须指定存放元素的类型和容量
	// 数组的长度是数组类型的一部分
	// var 数组变量名 [元素数量]T  T=>元素类型
	// a1，a2 是不同类型，数组间不能比较也不能赋值
	var a1 [3]bool // [false false false]
	var a2 [4]bool // [false false false false]
	fmt.Printf("a1:%T\n,a2:%T\n", a1, a2)

	// 数组的初始化
	// 如果数组不初始化，默认元素都是零值
	// 布尔值是 false， 整型和浮点型是 0 ，字符串是空字符串

	// 初始化方式1
	a1 = [3]bool{true, true, true}
	// 初始化方式2，根据初始值自动推动数组长度
	var a10 = [...]int{0, 1, 2, 3, 4, 5}
	fmt.Printf("a10的类型:%T\n", a10)
	fmt.Println("a10:", a10)
	// 初始化方式3, 根据索引初始化
	a3 := [5]int{0: 1, 4: 2}
	fmt.Println("a3:", a3) // [1 0 0 0 2]

	//数组的遍历
	fmt.Println("######数组的遍历#######")
	cites := [...]string{"北京", "上海", "深圳"}
	for i := 0; i < len(cites); i++ {
		fmt.Println(cites[i])
	}
	for index, value := range cites {
		fmt.Println(index, value)
	}
	fmt.Println("######多维数组#######")
	//多维数组
	var all [3][2]int
	all = [3][2]int{
		[2]int{1, 2},
		[2]int{3, 4},
		[2]int{5, 6},
	}
	fmt.Println("all:", all)
	//多维数组的遍历
	for i, v1 := range all {
		fmt.Println("第", i+1, "层v1:", v1)
		for _, v2 := range v1 {
			fmt.Println("v2:", v2)
		}
	}
	fmt.Println("######数组求和#######")
	// 数组求和
	var a13 = [...]int{1, 3, 5, 7, 8}
	var b int
	for _, v3 := range a13 {
		b += v3
	}
	fmt.Println(b)

	fmt.Println("########数组是值类型#########")
	// 数组是值类型
	// 赋值和传参会复制整个数组，因此会改变副本的值，不会改变本身的值
	b1 := [3]int{1, 2, 3}
	b2 := b1
	b2[0] = 100
	fmt.Println("b1:", b1) // [1 2 3]
	fmt.Println("b2:", b2) // [100 2 3]
}
