package main

import (
	"LearnGolang/package相关/calc"
	"fmt"
)

/*

包（package）是多个Go源码的集合，是一种高级的代码复用方案，Go语言为我们提供了很多内置包，如fmt、os、io等。

定义包
我们还可以根据自己的需要创建自己的包。一个包可以简单理解为一个存放.go文件的文件夹。 该文件夹下面的所有go文件都要在代码的第一行添加如下代码，声明该文件归属的包。

package 包名
注意事项：

一个文件夹下面直接包含的文件只能归属一个package，同样一个package的文件不能在多个文件夹下。
包名可以不和文件夹的名字一样，包名不能包含 - 符号。
包名为main的包为应用程序的入口包，这种包编译后会得到一个可执行文件，而编译不包含main包的源代码则不会得到可执行文件。

如果想在一个包中引用另外一个包里的标识符（如变量、常量、类型、函数等）时，该标识符必须是对外可见的（public）。
在Go语言中只需要将标识符的首字母大写就可以让标识符对外可见了。

在导入包名的时候，我们还可以为导入的包设置别名。通常用于导入的包名太长或者导入的包名冲突的情况。具体语法格式如下：

import 别名 "包的路径"

注意事项：
import导入语句通常放在文件开头包声明语句的下面。
导入的包名需要使用双引号包裹起来。
包名是从$GOPATH/src/后开始计算的，使用/进行路径分隔。
Go语言中禁止循环导入包。

*/

func main() {
	ret := calc.Add(2, 3)
	fmt.Println(ret)
}
