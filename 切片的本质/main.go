package main

import "fmt"

// 切片的本质就是一个框，框柱了一块连续的内存
// 切片属于引用类型，真正的数据都是保存在底层数组里
// make()函数创建切片
func main() {
	// 第二个参数是长度
	// 第三个参数是容量，不指定容量则等于长度
	// int类型没指定默认值 为0
	s1 := make([]int, 5)
	s2 := make([]int, 10)
	fmt.Printf("s1=%v len(s1)=%d cap(s1)=%d\n", s1, len(s1), cap(s1))
	// => s1=[0 0 0 0 0] len(s1)=5 cap(s1)=5
	fmt.Printf("s2=%v len(s2)=%d cap(s1)=%d\n", s2, len(s2), cap(s2))
	// => s2=[0 0 0 0 0 0 0 0 0 0] len(s2)=10 cap(s1)=10

	// 切片的赋值
	s3 := []int{1, 3, 5}
	s4 := s3
	fmt.Println(s3, s4)
	s3[0] = 1000 // 指向同一个数组,修改会造成s4的值也改变
	fmt.Println(s3, s4)
	// => [1000 3 5] [1000 3 5]

	// 切片的遍历
	// 1、索引遍历
	for i := 0; i < len(s3); i++ {
		fmt.Println(s3[i])
	}
	// 2.for range 循环
	for i, v := range s3 {
		fmt.Println(i, v)
	}
}
