package main

import "fmt"

func main() {
	age := 19
	if age > 18 {
		fmt.Println("澳门首家线上赌场开业啦！")
	} else {
		fmt.Println("暑假作业写完了吗？")
	}

	// 多个判断条件
	if age > 35 {
		fmt.Println("人到中年")
	} else if age > 18 {
		fmt.Println("青年")
	} else {
		fmt.Println("好好学习")
	}

	// 作用域
	// age1 变量 只在 if 语句中生效
	if age1 := 19; age1 > 18 {
		fmt.Println("澳门首家线上赌场开业啦！")
	} else {
		fmt.Println("暑假作业写完了吗？")
	}

	// go 语言中只有 for 循环
	// 可以使用 break return panic 语句强制退出循环
	// 基本格式
	for i := 0; i < 10; i++ {
		fmt.Println("基本格式：", i) // 0 1 2 3 4 5 6 7 8 9
	}

	// 变种1
	i := 5
	for ; i < 10; i++ {
		fmt.Println("变种1：", i) // 5 6 7 8 9
	}

	// 变种2
	n := 5
	for n < 10 {
		fmt.Println("变种2", n) // 5 6 7 8 9
		n++
	}

	// 死循环
	//for {
	//	fmt.Println("死循环了，重启吧！")
	//}

	// go语言中可以使用 for range 遍历 数组，切片，字符串， map结构， 及通道，通过 for range 遍历的返回值 有以下规律：
	// 1.数组 切片，字符串 返回索引和值
	// 2.map返回键和值
	// 3.通道只返回通道内的值

	s := "hello沙河"
	for i, v := range s {
		fmt.Printf("%d： %c\n", i, v)
	}
}
