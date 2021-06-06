package main

import "fmt"

// 切片:一个引用类型，都指向了底层的一个数组
func main() {
	//切片的定义
	var s1 []int    // 定义一个存放int类型的切片
	var s2 []string // 定义一个存放string类型的切片
	//初始化
	s1 = []int{1, 2, 3}
	s2 = []string{"const", "let", "var", "key"}
	fmt.Println(s1, s2)
	// len() 长度  cap() 容量
	fmt.Printf("len(s1): %d cap(s1):%d\n", len(s1), cap(s1))
	fmt.Printf("len(s2): %d cap(s2):%d\n", len(s2), cap(s2))
	// 由数组得到切片
	a1 := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	s3 := a1[0:4]          // 基于数组去切割，左包含右不包含（左闭右开）
	s4 := a1[:4]           // 从零切到4 => [0:4]
	s5 := a1[2:]           // 从2切到末尾 => [2:len(ai)]
	s6 := a1[:]            // 从头到尾 => [0: len(a1)]
	fmt.Println("s3:", s3) // => [1 2 3 4]
	fmt.Println("s4:", s4) // => [1 2 3 4]
	fmt.Println("s5:", s5) // => [3 4 5 6 7 8 9]
	fmt.Println("s6:", s6) // => [1 2 3 4 5 6 7 8 9]
	// 切片的容量是指底层数组的容量（底层数组从切片的第一个元素到最后的元素的数量）
	// 切片的长度就是元素的个数
	fmt.Printf("len(s4):%d cap(s4):%d\n", len(s4), cap(s4)) // => len(s4):4 cap(s4):9
	fmt.Printf("len(s5):%d cap(s5):%d\n", len(s5), cap(s5)) // => len(s5):7 cap(s5):7
	// 切片再切割
	s7 := s5[3:]
	fmt.Println(s7)                                         // => [6 7 8 9]
	fmt.Printf("len(s7):%d cap(s7):%d\n", len(s7), cap(s7)) // => len(s7):4 cap(s7):4
	//切片是引用类型， 底层都指向了一个数组，修改数组，切片的值也会修改
	a1[3] = 400
	a1[7] = 800
	fmt.Println("s3:", s3) // => [1 2 3 400]
	fmt.Println("s4:", s4) // => [1 2 3 400]
	fmt.Println("s5:", s5) // => [3 4 5 6 7 800 9]
	fmt.Println("s6:", s6) // => [1 2 3 4 5 6 7 800 9]
}
