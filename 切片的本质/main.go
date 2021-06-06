package main

import "fmt"

// 切片的本质就是一个框，框柱了一块连续的内存
// 切片属于引用类型，真正的数据都是保存在底层数组里
// make()函数创建切片
func main() {
	// 第二个参数是长度
	// 第三个参数是容量，不指定容量则等于长度
	// int类型没指定默认值 为0
	s := make([]int, 5, 10)
	s1 := make([]int, 5)
	s2 := make([]int, 10)
	fmt.Printf("s1=%v len(s1)=%d cap(s1)=%d\n", s, len(s), cap(s))    // =>[0 0 0 0 0] 5 10
	fmt.Printf("s1=%v len(s1)=%d cap(s1)=%d\n", s1, len(s1), cap(s1)) // =>[0 0 0 0 0] 5 5
	fmt.Printf("s2=%v len(s2)=%d cap(s1)=%d\n", s2, len(s2), cap(s2)) // =>[0 0 0 0 0 0 0 0 0 0] 10 10

	fmt.Println("####切片不能直接比较####")
	// 切片是不能直接比较的，不能用 ==  操作符来判断两个切片时候含有全部相等的元素
	// 切片唯一合法的比较操作是和 nil（空） 比较， 一个 nil 值的切片没有底层数组，长度和容量都是0
	// 但我们不能说一个长度和容量都是0的切片一定是 nil
	// 要切片是不是空的，要用 len（s） == 0 来判断，不应该使用 nil
	var n1 []int
	n2 := []int{}
	n3 := make([]int, 0)
	fmt.Println(n1 == nil, n2 == nil, n3 == nil) // true false false

	fmt.Println("####切片的赋值####")
	// 切片的赋值
	s3 := []int{1, 3, 5}
	s4 := s3
	fmt.Println("改变前：s3", s3, "s4", s4)
	s3[0] = 1000 // 指向同一个数组,修改会造成s4的值也改变
	fmt.Println("改变后：s3", s3, "s4", s4)
	// => [1000 3 5] [1000 3 5]

	fmt.Println("####切片的遍历####")
	// 切片的遍历
	// 1、索引遍历
	fmt.Println("----索引遍历----")
	for i := 0; i < len(s3); i++ {
		fmt.Println(s3[i])
	}
	fmt.Println("----for range 循环----")
	// 2.for range 循环
	for i, v := range s3 {
		fmt.Println(i, v)
	}

}
