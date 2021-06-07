package main

import "fmt"

func main() {
	// go 语言中不存在指针操作，只需要记住两个操作
	// &  =》 取地址
	// *  =》 根据地址取值

	n := 18
	p := &n
	fmt.Println(p)
	fmt.Printf("%T\n", p)
	// 根据地址取值
	m := *p
	fmt.Println(m)
	fmt.Printf("%T\n", m)

	//	地址操作符&和取值操作符*是一对互补操作符，&取出地址，*根据地址取出地址指向的值。
	//	变量、指针地址、指针变量、取地址、取值的相互关系和特性如下：
	//	对变量进行取地址（&）操作，可以获得这个变量的指针变量。
	//	指针变量的值是指针地址。
	//	对指针变量进行取值（*）操作，可以获得指针变量指向的原变量的值。

	//	fmt.Println("#### new ####")
	//	var a *int  // 这里其实是个空指针， 零 值 nil
	//	*a = 100
	//	fmt.Println(*a)
	//	报错 panic: runtime error: invalid memory address or nil pointer dereference

	// new 函数申请一个内存地址
	var a = new(int)
	*a = 100
	fmt.Println("a:", a)

	//	make也是用于内存分配的，区别于new，它只用于slice、map以及chan的内存创建，而且它返回的类型就是这三个类型本身，而不是他们的指针类型，因为这三种类型就是引用类型，所以就没有必要返回他们的指针了。
	//	make函数是无可替代的，我们在使用slice、map以及channel的时候，都需要使用make进行初始化，然后才可以对它们进行操作

	//new与make的区别
	//二者都是用来做内存分配的。
	//make只用于slice、map以及channel的初始化，返回的还是这三个引用类型本身；
	//而new用于类型的内存分配，并且内存对应的值为类型零值，返回的是指向类型的指针。
}
