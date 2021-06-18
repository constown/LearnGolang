package main

import "fmt"

// 方法和接收者

type dog struct {
	name string
}

// 构造函数
func newDog(name string) dog {
	return dog{
		name: name,
	}
}

// 方法是作用于特定类型的函数
// 接受者表示的是调用该方法的具体类型变量，多用于类型名首字母小写表示
// 值接收者，也可以接收指针
// 什么时候用指针类型接受者？
// 1、需要修改接收者中的值
// 2、接收者是拷贝代价比较大的对象
// 3、保证一致性，如果有某个方法使用你指针接收者，那么其他地方也应该使用指针接收者
func (d dog) wow() {
	fmt.Printf("%s: 汪汪汪\n", d.name)
}

func (d *dog) w() {
	fmt.Printf("%s: 汪汪汪", d.name)
}

// 标识符：变量名 函数名 类型名 方法名
// go语言中如果标识符首字母大写，就表示对外部可见（暴露的，公共的），就要写注释

// Dog 这是一个狗的结构体， 这里要写注释
type Dog struct {
	name string
}

// 给自定义类型添加方法
// 不能给别的包的类型添加方法，只能给自己的包里面的类型添加方法

type myInt int

func (m myInt) add() {
	fmt.Println("+++")
}

func main() {
	d1 := newDog("jack")
	d1.wow()
	m := myInt(100)
	m.add()
}
