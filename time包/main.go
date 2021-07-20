package main

import (
	"fmt"
	"time"
)

// 时间格式化 2006-01-02 15：04：05.000

// 时间类型 time.Time
// 时间戳 time.now().Unix() 1970.1.1到现在的秒数
// 		 time.now().UnixNano() 纳秒  1970.1.1到现在的纳秒数

// 时间间隔类型 time.Duration  表示两个时间经过的秒数
//			  time.Second
//			  time.Hour   等等
//	const (
//		Nanosecond  Duration = 1
//		Microsecond          = 1000 * Nanosecond
//		Millisecond          = 1000 * Microsecond
//		Second               = 1000 * Millisecond
//		Minute               = 60 * Second
//		Hour                 = 60 * Minute
//	)

// 时间操作
// 时间对象 +/- 一个时间间隔对象

//  定时器
//  time.Tick(time.Second)

// 解析字符串格式时间（时区）按照某一个格式去解析，传过来的字符串
// time.Parse("2006-01-02 15:04:05", "2021-07-06 23:18:50") 解析成一个utc时间
// 如果要解析成时区的时间，需要定义一个时区，拿到时区对象后 time.ParseInLocation 解析成时区时间

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
	fmt.Println(timeObj, "timeObj")
	// sub 两个时间相减
	td := timeObj.Sub(now)
	fmt.Println(td)
}
