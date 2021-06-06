package main

import "fmt"

func main() {
	// 当 i= 5 的时候跳出循环
	for i := 0; i < 10; i++ {
		if i == 5 {
			break
		}
		fmt.Println(i)
	}
	fmt.Println("over")
	fmt.Println("#############")
	// 当 i = 5 的时候，跳过此次循环，继续下一次循环
	for i := 0; i < 10; i++ {
		if i == 5 {
			continue
		}
		fmt.Println(i)
	}
	fmt.Println("over")

	// switch 简化大量的判断
	fmt.Println("#############")
	n := 5
	switch n {
	case 1:
		fmt.Println("大拇指")
	case 2:
		fmt.Println("食指")
	case 3:
		fmt.Println("中指")
	case 4:
		fmt.Println("无名指")
	case 5:
		fmt.Println("小拇指")
	default:
		fmt.Println("无效的数字")
	}
	fmt.Println("#############")
	switch b := 7; b {
	case 1, 3, 5, 7, 9:
		fmt.Println("奇数")
	case 2, 4, 6, 8:
		fmt.Println("偶数")
	default:
		fmt.Println("不支持的数字")
	}
	fmt.Println("#############")
	//goto 跳出多层 for循环
	for i := 0; i < 10; i++ {
		for j := 'A'; j < 'Z'; j++ {
			if j == 'C' {
				break
			}
			fmt.Printf("i:%v-j:%c\n", i, j)
		}
	}
	fmt.Println("#############")
	flag := true
	for i := 0; i < 10; i++ {
		for j := 'A'; j < 'Z'; j++ {
			if j == 'C' {
				break
			}
			fmt.Printf("i:%v-j:%c\n", i, j)
		}
		if flag {
			break
		}
	}
	fmt.Println("#############")
	// 使用goto
	for i := 0; i < 10; i++ {
		for j := 'A'; j < 'Z'; j++ {
			if j == 'C' {
				goto breakFunc // 跳转到指定的那个标签
			}
			fmt.Printf("i:%v-j:%c\n", i, j)
		}
	}
breakFunc: // label 标签
	fmt.Println("直接跳出了循环")
}
