package main

import "fmt"

func main() {
	// 算数运算符：
	// 加 +
	// 减 -
	// 乘 *
	// 除 /
	// 求余 %
	a, b := 5, 2
	fmt.Println(a + b) // 7
	fmt.Println(a - b) // 3
	fmt.Println(a * b) // 10
	fmt.Println(a / b) // 2
	fmt.Println(a % b) // 1
	fmt.Println("#######")
	// ++ -- 都是单独的语句 不恩呢放在 =  的右边赋值 代表 +1 -1
	a++
	b--
	fmt.Println(a)
	fmt.Println(b)

	// 关系运算符
	// == 判断相等 go语言是强类型的，相同类型的变量才能比较
	// != 不等于
	// >= 大于等于
	// <= 小于等于
	// > 大于
	// < 小于

	fmt.Println(a == b)
	fmt.Println(a != b)
	fmt.Println(a < b)
	fmt.Println(a <= b)
	fmt.Println(a > b)
	fmt.Println(a >= b)

	fmt.Println("&&:###########")
	// 逻辑运算符
	// && 且， 都为真才为真，有一个假就为假
	// || 或， 只要有一个为真就为真，都为假才为假
	age := 22
	// 年龄大于 18 且 小于 60
	if age > 18 && age < 60 {
		fmt.Println("社畜")
	} else {
		fmt.Println("不上班")
	}
	fmt.Println("||:###########")
	// n年龄小于18 或 年龄 大于 60
	if age < 18 || age > 60 {
		fmt.Println("不用上班")
	} else {
		fmt.Println("社畜")
	}

	// ! 取反

	// 位运算符: 针对2进制数
	// & 按位与 参与预算的两数各对应的二进位相与, 都为1才为1
	// 5: 101
	// 2:  10
	//    000 => 0
	fmt.Println("按位与:#########")
	fmt.Println(5 & 2)
	// | 按位或 参与预算的两数各对应的二进位相与, 有一个为1就为1
	// 5: 101
	// 2:  10
	//    111 => 7
	fmt.Println("按位或:#########")
	fmt.Println(5 | 2)

	// ^ 按位异或 参与预算的两数各对应的二进位相与, 两位不一样就位1
	// 5: 101
	// 2:  10
	//    111 => 7
	fmt.Println("按位异或:#########")
	fmt.Println(5 ^ 2)

	// << 左移 左移 n 位就是乘以 2 的 n 次方
	fmt.Println("左移:#########")
	// 5: 101
	//   1010 => 10
	fmt.Println(5 << 1)

	// >> 右移 右移 n 位就是乘以 2 的 n 次方
	// 5: 101
	//	   10 => 2
	fmt.Println("右移:#########")
	fmt.Println(5 >> 1)
	m := int8(1)
	fmt.Println(m << 10) // 只能装8位，早就值展示最后 8 位
}
