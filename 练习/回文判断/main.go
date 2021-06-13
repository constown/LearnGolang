package main

import "fmt"

func main() {
	//  回文判断
	//	字符串从左往右和从右往左读是一样的，就是回文
	//	上海自来水来自海上
	//	山西运煤车煤运西山
	//	黄山落叶松叶落山黄
	ss := "上海自来水来自海上"

	// 	把字符串中的字符拿出来放到一个 []rune 中
	r := make([]rune, 0, len(ss))
	for _, c := range ss {
		r = append(r, c)
	}

	for i := 0; i < len(r)/2; i++ {
		if r[i] != r[len(r)-1-i] {
			fmt.Println("不是回文")
			return
		}
	}
	fmt.Println("是回文")
}
