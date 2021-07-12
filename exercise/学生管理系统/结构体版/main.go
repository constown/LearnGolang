package main

import (
	"fmt"
	"os"
)

// 声明一个全局的变量 学生管理对象
var manage = studentManage{
	allStudent: make(map[int64]student, 100),
}

// 菜单函数
func showMenu() {
	fmt.Println("welcome sms!")
	fmt.Println(`
	1.查看所有学生
	2.添加学生
	3.修改学生
	4.删除学生
	5.退出
	`)
	fmt.Print("请输入要执行的操作:")
}

func main() {
	for {
		showMenu()
		// 等待用户输入
		var choice int
		fmt.Scanln(&choice)
		fmt.Println("您输入的是:", choice)
		// 执行对应的操作
		switch choice {
		case 1:
			manage.showStudent()
		case 2:
			manage.addStudent()
		case 3:
			manage.editStudent()
		case 4:
			manage.deleteStudent()
		case 5:
			os.Exit(1)
		default:
			fmt.Println("无效的操作")
		}
	}
}
