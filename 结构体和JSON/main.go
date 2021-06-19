package main

import (
	"encoding/json"
	"fmt"
)

// 结构体和JSON
// 1.把go语言中的结构体变量 ==> JSON格式的字符串    	序列化
// 2.把JSON格式的字符串 ==> go语言中的结构体变量		反序列化

type person struct {
	Name string `json:"name" db:"name" ini:"name"` // 可以添加tag,代表不同的地方叫啥名字,比如这个在json里面叫name
	Age  int    `json:"age"`                       // 这里首字母大写是因为 在 json这个包里去拿对应的东西,所以在别的包里去拿,小写是拿不到的,必须要大写成暴露的
}

func main() {
	p1 := person{
		Name: "周玲",
		Age:  18,
	}
	// 序列化使用 json.Marshal
	b, err := json.Marshal(p1)
	if err != nil {
		fmt.Printf("marshal failed, err:%v", err)
		return
	} else {
		fmt.Printf("%#v\n", string(b))
	}

	// 反序列化
	str := `{"name":"lucy","age":18}`
	var p2 person
	json.Unmarshal([]byte(str), &p2)
	fmt.Printf("%#v\n", p2)
}
