package main

import (
	"fmt"
	"strconv"
)

// strconv

func main() {
	str := "1000"
	//ret1 := int64(str)
	ret1, err := strconv.ParseInt(str, 10, 64) // 转换成10进制 64位的 int   返回值是int64
	if err != nil {
		fmt.Println("parseint failed,err:", err)
		return
	}
	fmt.Printf("%#v %T\n", ret1, ret1)

	// 字符串 和 int 转换
	//strconv.Atoi()  strconv.Itoa
	retInt, _ := strconv.Atoi(str)
	fmt.Printf("%#v %T\n", retInt, retInt)

	i := int32(100)
	// 不能直接用string转，这个等于去找编码了
	//ret2 := string(i)
	//fmt.Println(ret1)
	ret2 := fmt.Sprintf("%d", i)
	fmt.Printf("%#v\n", ret2)

	// 从字符串中解析出布尔值
	boolStr := "false"
	boolValue, _ := strconv.ParseBool(boolStr)
	fmt.Printf("%#v %T\n", boolValue, boolValue)
}
