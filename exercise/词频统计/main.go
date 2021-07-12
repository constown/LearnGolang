package main

import (
	"fmt"
	"strings"
)

func main() {
	//	判断单词出现的次数
	// 	把字符串切割成切片
	//	遍历切片存储到一个map
	str := "how do you do"
	count := make(map[string]int)
	str2 := strings.Split(str, " ")
	for _, value := range str2 {
		count[value]++
	}
	fmt.Println(count)
}
