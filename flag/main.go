package main

import (
	"flag"
	"fmt"
	"os"
	"time"
)

//os.Args 获取命令行参数
func f1() {
	fmt.Printf("%#v\n", os.Args)
	fmt.Printf("%T\n", os.Args)
}

// flag 获取命令行参数
func main() {
	f1()
	var hobby string
	// 创建一个标志位参数
	flag.StringVar(&hobby, "hobby", "hobby", "hobby")
	name := flag.String("name", "丸子", "请输入名字")
	age := flag.Int("age", 18, "请输入年龄")
	married := flag.Bool("married", false, "结婚了吗")
	marriedTime := flag.Duration("marriedTime", time.Second, "结婚了吗")
	//	使用flag
	flag.Parse()
	fmt.Println(*name, *age, *married, *marriedTime, hobby)
	//	flag.Args() 返回命令行参数后面的其他参数， []string 类型
	//	flag.NArg() 返回命令行参数后的其他参数个数
	//	flag.NFlag() 返回命令行参数的使用个数
}
