package main

import "fmt"

func main() {
	var a1 [3]bool
	var a2 [4]bool
	fmt.Printf("a1:%T,a2:%T\n", a1, a2)
	var a10 = [...]int{0, 1, 2, 3, 4, 5}
	fmt.Printf("a10:%T", a10)
	fmt.Println(a10)

	//数组的遍历
	cites := [...]string{"备机", "伤害", "设置"}
	for i := 0; i < len(cites); i++ {
		fmt.Println(cites[i])
	}
	for i, v := range cites {
		fmt.Println(i, v)
	}
	//多维数组
	var all [3][2]int
	all = [3][2]int{
		[2]int{1, 2},
		[2]int{3, 4},
		[2]int{5, 6},
	}
	fmt.Println(all)
	//多维数组的遍历
	for _, v1 := range all {
		fmt.Println(v1)
		for _, v2 := range v1 {
			fmt.Println(v2)
		}
	}
	// 数组求和
	var a13 = [...]int{1, 3, 5, 7, 8}
	var b int
	for _, v3 := range a13 {
		b += v3
	}
	fmt.Println(b)
}
