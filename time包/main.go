package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Println(now)
	fmt.Println(now.UnixNano())
	fmt.Println(now.Second())
	fmt.Println(now.Unix())
	fmt.Println(now.Minute())
	fmt.Println(now.Hour())
	fmt.Println(now.Day())
	fmt.Println(now.Month())
	fmt.Println(now.Year())
	fmt.Println(now.Format("2006-01-02"))
	// 时间解析
	timeObj, err := time.Parse("2006-01-02", "2021-03-02")
	if err != nil {
		fmt.Println("err: ", err)
		return
	}
	fmt.Println(timeObj)
}
