package main

import "fmt"

// 空接口 可以接收任何类型的值 interface{}  没有必要起名字
// 所有的类型都实现了空接口,也就是任意类型的变量都能保存在空接口

// 空接口作为函数参数
func show(a interface{}) {
	fmt.Printf("type:%T value:%v\n", a, a)
}

// 类型断言
func assert(a interface{}) {
	fmt.Printf("%T\n", a)
	str, ok := a.(string) // 类型断言  这里接受两个参数
	if !ok {
		fmt.Println("猜错了")
	} else {
		fmt.Printf("传进来的是一个字符串")
	}
	fmt.Println(str)
}

func assert2(a interface{}) {
	fmt.Printf("%T\n", a)
	switch t := a.(type) {
	case string:
		fmt.Println("是一个字符串:", t)
	case int:
		fmt.Println("是一个int:", t)
	case bool:
		fmt.Println("是一个bool:", t)
	case int64:
		fmt.Println("是一个int64:", t)
	}
}

func main() {
	m1 := make(map[string]interface{}, 16)
	m1["name"] = "周玲"
	m1["age"] = 18
	m1["marry"] = false
	m1["hobby"] = [...]string{"唱", "跳", "rap"}
	fmt.Println(m1)

	show(false)
	show(nil)
	show(m1)

	assert(100)
	assert2(false)
}
