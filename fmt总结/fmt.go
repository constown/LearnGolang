package main

import "fmt"

// fmt 主要分为向外输出内容和获取输入内容两大部分
// 主要提供了以下几种输出相关函数

// Print
// fmt.Print  	直接输出
// fmt.Println  带换行
// fmt.Printf	带格式化

// 获取输入
// fmt.Scan
// fmt.Scanf
// fmt.Scanln

func main() {
	fmt.Print("沙河")
	fmt.Println("娜扎")
	str := "牛蛙"
	fmt.Printf("%s\n", str)
	// 整数 => 字符
	fmt.Printf("%q\n", 65)

	// h获取用户的输入
	var s string
	fmt.Scan(&s)
	fmt.Println("用户输入的值：", s)
	var (
		name  string
		age   int
		class string
	)
	fmt.Scanf("%s %d %s\n", &name, &age, &class)
	fmt.Println("name:", name, "age:", age, "class:", class)
	fmt.Scanln(&name, &age, &class)
	fmt.Println("name:", name, "age:", age, "class:", class)
}
