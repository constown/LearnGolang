package main

import (
	"fmt"
	"strings"
)

func main() {
	// go语言中 字符串只能用 双引号 包裹 “字符串”
	//	单引号包裹的是字符
	s := "hello"
	fmt.Println(s)
	// 单独的字母，房子，符号表示一个字符
	c1 := 'h'
	fmt.Println(c1)

	// 转义符
	//  \r  回车符
	//	\n	换行符
	//	\t	制表符
	//	\'	单引号
	//	\"	双引号
	//	\\	反斜杠
	fmt.Println("\"E:\\Golang\\src\\HelloGo\"")

	// 多行字符串

	s1 := `
	床前明月光，
	  疑是地上霜.
	举头望明月，
	  低头思故乡。
	`
	fmt.Println(s1)

	//	字符串的操作
	// len(str) 求长度
	s2 := "love"
	sLen := len(s2)
	fmt.Printf("长度: %v\n", sLen)

	// + 或 fmt.Sprintf 拼接字符串
	firstName := "司马"
	lastName := "噫"
	fullName1 := firstName + lastName
	fullName2 := fmt.Sprintf("%s%s\n", firstName, lastName)
	fmt.Printf("字符串拼接: %v\n", fullName1)
	fmt.Printf("字符串拼接: %v\n", fullName2)

	//	string.Split 分割
	s3 := `E:\Golang\src\HelloGo`
	ret := strings.Split(s3, "\\") // [E: Golang src HelloGo]
	fmt.Println(ret)

	//  string.contains 判断是否包含
	fmt.Println(strings.Contains(s3, "src")) // true

	//	string.HasPrefix 前缀判断
	fmt.Println(strings.HasPrefix(s3, "D:")) // false

	//  string.HasSuffix 后缀判断
	fmt.Println(strings.HasSuffix(s3, "Go")) // true

	//  string.Index(), string.LastIndex() 子串出现的位置 (第一次出现， 最后一次出现)
	s4 := "abcdedcba"
	fmt.Println(strings.Index(s4, "c"))     // 2
	fmt.Println(strings.LastIndex(s4, "b")) // 7

	//  string.Join(a[]string, sep string) join操作  数组拼成字符串
	fmt.Println(strings.Join(ret, "+"))
}
