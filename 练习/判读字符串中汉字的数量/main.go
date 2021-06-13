package main

import (
	"fmt"
	"unicode"
)

func main() {
	// 1.判断字符串中汉字的数量
	//	难点：判断一个字符是汉字
	s1 := "hello沙河"
	//	1.依次拿到字符串的字符
	var count int
	for _, c := range s1 {
		//	2.判断是不是汉字
		if unicode.Is(unicode.Han, c) {
			count++
		}
	}
	// 	3.把汉字出现的次数累计
	fmt.Println(count)
}
