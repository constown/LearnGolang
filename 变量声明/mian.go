package main

import "fmt"

// main 函数是包的入口

// 函数外只能放置标识符（变量、常量。函数，类型声明）， 语句不能放在函数外

// 声明变量
// go预言中的变量必须先声明在使用
// 使用 var 关键字声明变量
// var s1 string => var 变量名 变量数据类型
// go语言中推荐使用小驼峰命名变量 studentName

// 单独声明
//var name string
//var age int
//var isOK bool

//批量声明
var (
	name string
	age  int
	isOK bool
)

// main 是程序的入口函数

func main() {
	name = "理想"
	age = 16
	isOK = false

	// go语言中的变量声明了必须使用，否则无法编译
	fmt.Print(isOK)              //
	fmt.Printf("name: %s", name) // %s 占位符，使用name变量值去替换占位符
	fmt.Println(age)             // %d 十进制 int 类型

	// print 家族：
	//fmt.Print()  打印一句话
	//fmt.Printf() 带有格式化的
	//fmt.Println() 带有一个换行符 打印完内容后自动加一个换行
}
