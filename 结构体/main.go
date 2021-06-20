package main

import (
	"fmt"
)

// 结构体
// 可以理解为 js中的对象？
// 结构体里的字段是唯一的
// 结构体是一种数据类型,一种我们可以自己造的保存多个维度数据的类型
// 结构体是值类型  [ 注意: 值类型和引用类型的区别 ]  要修改值类型需要使用指针

// 使用 [ type 结构体名字 struct ]来进行声明

type person struct {
	name   string
	age    int
	gender string
	hobby  []string
}

// 结构体的匿名字段 把类型作为字段名了，相同类型就只能写一个了
// 只适用于字段简单且比较少的，不常用
type personA struct {
	string
	int
}

type address struct {
	province string
	city     string
}

// 结构体的嵌套  把另一个结构体嵌套到另一个结构体
type personB struct {
	name    string
	age     int
	address address
}

type company struct {
	name    string
	address address
}

// 匿名嵌套结构体
type personC struct {
	name string
	address
}

func main() {
	var lucy person
	// 通过字段来赋值
	lucy.name = "lucy"
	lucy.age = 18
	lucy.gender = "女"
	lucy.hobby = []string{"足球", "篮球", "双色球"}
	fmt.Printf("类型：%T value：%v\n", lucy, lucy)

	// 匿名结构体 多用于临时场景
	var s struct {
		x string
		y int
	}

	s.x = "嘿嘿"
	s.y = 180
	fmt.Printf("类型：%T value：%v\n", s, s)

	// 结构体的匿名字段
	p1 := personA{
		"周宁",
		900,
	}
	fmt.Println(p1.string)

	p2 := personB{
		name: "周宁",
		age:  90,
		address: address{
			province: "山东",
			city:     "威海",
		},
	}
	fmt.Println(p2)
	fmt.Println(p2.address.city)

	p3 := personC{
		name: "周宁",
		address: address{
			province: "山东",
			city:     "威海",
		},
	}
	fmt.Println(p3)
	fmt.Println(p3.city) // 使用匿名嵌套结构体的时候，可以用这种语法糖

	// 匿名结构体
	var ss = struct {
		x int
		y int
	}{10, 20}
	fmt.Println("ss:", ss)
}
