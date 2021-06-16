package main

import "fmt"

// panic 和 recover
// recover 必须搭配defer使用
// defer一定要在可能引发 panic 的语句之前定义

func funcA() {
	fmt.Println("程序开始执行")
}
func funcB() {
	// 刚刚打开数据库链接
	// defer语句可以用来释放资源
	defer func() {
		err := recover()
		fmt.Println("err:", err)
		fmt.Println("释放数据库链接")
	}()
	panic("数据库链接失败，出现了严重错误") // 程序崩溃推出
	fmt.Println("程序正常执行")
}
func funcC() {
	fmt.Println("程序执行完成")

}
func main() {
	funcA()
	funcB()
	funcC()
}
