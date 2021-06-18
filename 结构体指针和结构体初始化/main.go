package main

import "fmt"

// 结构体是值类型

type person struct {
	name, gender string // 字段的类型简写
}

// go语言中函数参数永远是拷贝
func f(x person) {
	x.gender = "女" // 这里修改的副本 内部不会影响到外部的值
}

// 修改的是内存地址，等于改的就是原来的数据
func f2(x *person) {
	(*x).gender = "女" // 这是根据内存地址找到原变量，修改的是原来的那个变量
	// 语法糖
	x.gender = "未知" // 自动根据指针找对应的变量
}

func main() {
	var p person
	p.name = "周宁"
	p.gender = "男"
	f(p)
	fmt.Println(p.gender) // 男
	f2(&p)
	fmt.Println(p.gender) // 女

	// 创建指针类型的结构体
	var p2 = new(person)
	p2.name = "jack"
	fmt.Printf("%T\n", p2) // 返回的是指针类型 *person
	fmt.Printf("%v\n", p2)

	// new 和 make 的区别？
	// new是基本类型也包括值类型返回的值是对应的指针  make是特定类型，返回的是值

	// 结构体指针2
	// key value 初始化
	var p3 = &person{
		name:   "远方",
		gender: "男",
	}
	fmt.Println(p3)

	// 使用值列表的形式初始化，值的顺序必须和结构体定义时字段的顺序一致
	p4 := &person{
		"小蜗牛",
		"女",
	}
	fmt.Println(p4)

	// 结构体内存布局
	// 一个结构里面的内存是连续的
}
