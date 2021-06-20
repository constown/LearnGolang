package main

import "fmt"

/*
Go语言提倡面向接口编程。
每个接口由数个方法组成，接口的定义格式如下：
type 接口类型名 interface{
    方法名1( 参数列表1 ) 返回值列表1
    方法名2( 参数列表2 ) 返回值列表2
    …
}
其中：
接口名：使用type将接口定义为自定义的类型名。Go语言的接口在命名时，一般会在单词后面添加er，如有写操作的接口叫Writer，有字符串功能的接口叫Stringer等。接口名最好要能突出该接口的类型含义。
方法名：当方法名首字母是大写且这个接口类型名首字母也是大写时，这个方法可以被接口所在的包（package）之外的代码访问。
参数列表、返回值列表：参数列表和返回值列表中的参数变量名可以省略。
举个例子：
type writer interface{
    Write([]byte) error
}
当你看到这个接口类型的值时，你不知道它是什么，唯一知道的就是可以通过它的Write方法来做一些事情。
*/

// 结构的实现

type animal interface {
	move()
	eat(something string)
}

type cat struct {
	name string
	feet int8
}

func (c cat) move() {
	fmt.Println("cat走猫步")
}

func (c cat) eat() {
	fmt.Println("cat吃鱼")
}

type chicken struct {
	feet int8
}

func (c chicken) move() {
	fmt.Println("chicken母鸡")
}

// 实现的方法必须严格符合接口的要求,才算实现了接口
func (c chicken) eat(food string) {
	fmt.Printf("chicken吃%s\n", food)
}
func main() {
	var a1 animal   // 定义一个接口类型的变量
	blueCat := cat{ // 定义一个cat类型的变量
		name: "淘气",
		feet: 4,
	}
	// a1 = blueCat 报错  因为 eat 方法没有不符合接口的eat   没有参数
	fmt.Println(blueCat)

	hen := chicken{
		feet: 2,
	}

	a1 = hen // 接口是作为一种类型来使用的
	fmt.Println(a1)

	a1.eat("小黄鱼")
}
