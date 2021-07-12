package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now() // 本地的时间
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
	n := 100
	time.Sleep(time.Duration(n))

	// 明天的这个时间
	// 按照指定格式去解析一个字符串格式的时间
	time.Parse("2006-01-02 15:04:05", "2021-07-06 23:18:50")
	// 按照东八区的时区和格式去解析一个字符串格式的时间
	loc, err := time.LoadLocation("Asia/shanghai")
	if err != nil {
		fmt.Printf("load loc faild, err:%v\n", err)
	}
	// 按照指定时区解析时间
	timeObj, err = time.ParseInLocation("2006-01-02 15:04:05", "2021-07-06 23:18:50", loc)
	if err != nil {
		fmt.Printf("parse time failed, err:%v\n", err)
		return
	}
	fmt.Println(timeObj, "tiemObj")
	td := timeObj.Sub(now)
	fmt.Println(td)
}
