package main

import "fmt"

// append() 为切片追加元素 原来的底层数组放不下的时候，在底层数组换一个更大的新数组
func main() {
	s1 := []string{"北京", "上海", "深圳"}
	fmt.Printf("s1=%v len(s1)=%d cap(s1)=%d\n", s1, len(s1), cap(s1))
	// s1[3] = "广州" // 错误的写法，会导致编译错误，索引越界 =》 index out of range
	// 调用append（）必须用原来的切片变量接收返回值
	s1 = append(s1, "广州")
	fmt.Printf("s1=%v len(s1)=%d cap(s1)=%d\n", s1, len(s1), cap(s1))
	ss := []string{"杭州", "苏州"}
	s1 = append(s1, ss...) // ...可以看做JS的扩展运算符

	//	关于 append 删除切片中的某个元素
	a1 := [...]int{1, 3, 5, 7, 9, 11, 13, 15, 17}
	a2 := a1[:]

	//	删除索引为 1 的那个 3
	a2 = append(a2[:1], a2[2:]...)
	fmt.Println("a2:", a2)
	fmt.Println("a1:", a1)
}
