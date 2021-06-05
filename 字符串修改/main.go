package main

import "fmt"

func main() {
	s1 := "Hell沙河모래강"
	// len() 得到的是 byte 字节的长度
	//Go 语言的字符有以下两种：
	//uint8类型，或者叫 byte 型，代表了ASCII码的一个字符。
	//rune类型，代表一个 UTF-8字符。

	fmt.Println("len长度", len(s1)) // 19

	//for i := 0; i < len(s1); i++ {
	//	//fmt.Println(s1[i])
	//	fmt.Printf("%c\n", s1[i])
	//}

	// 取出具体的字符
	//for _, c := range s1 {
	//	//fmt.Println(c)
	//	fmt.Printf("%c\n", c) // %c 字符
	//}

	// 字符串不能直接修改

	s2 := "白萝卜"
	s3 := []rune(s2)        // 转换成 rune 切片 =>  [白 萝 卜]
	s3[0] = '红'             // 这里应该是字符 红
	fmt.Println(string(s3)) // 在强制转换为 string

	c1 := "红"
	c2 := '红'
	fmt.Printf("c1：%T c2:%T\n", c1, c2) // 一个是字符串 一个是字符  c1：string  c2:int32
	c3 := "H"
	c4 := 'H'
	fmt.Printf("c3：%T c4:%T\n", c3, c4) // 一个是字符串 一个是字符  c3：string  c4:int32
	fmt.Printf("%d\n", c4)              // 72

	// 类型转换
	n := 10
	var f float64
	f = float64(n)
	fmt.Println("f：", f)
}
