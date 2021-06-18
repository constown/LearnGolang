package main

import (
	"fmt"
	"os"
)

/*
	函数版学生管理系统
	写一个系统能够参看/新增/删除学生
*/

type student struct {
	id   uint64
	name string
}

var (
	allStudents map[uint64]*student // 变量声明
)

func newStudent(id uint64, name string) *student {
	return &student{
		id:   id,
		name: name,
	}
}

func showALlStudents() {
	// 打印所有的学生
	fmt.Println("⬇️⬇️⬇️ 查询到学生列表如下 ⬇️⬇️⬇️")
	for key, value := range allStudents {
		fmt.Printf("学号：%d 姓名：%s\n", key, value.name)
	}
	fmt.Println("----------查询结束-------------")
	fmt.Println("——————————————————————————————")
}
func addStudent() {
	// 向allStudents添加一个学生
	// 1、新建一个学生
	var (
		id   uint64
		name string
	)
	id = uint64(len(allStudents))
	fmt.Print("👌请输入学生的名字：")
	fmt.Scanln(&name)
	// 调用构造函数新建学生
	newStu := newStudent(id, name)
	// 2、追加进去
	allStudents[id] = newStu
	fmt.Println("🎉🎉🎉 创建成功！🎉🎉🎉")
	fmt.Println("——————————————————————————————")
}
func deleteStudent() {
	// 1、用户输入要删除的学号
	var (
		deleteId uint64
	)
	fmt.Print("👌请输入要删除的学生id：")
	fmt.Scanln(&deleteId)
	// 删除对应的键值对
	delete(allStudents, deleteId)
	fmt.Println("🎉🎉🎉 删除成功！🎉🎉🎉")
	fmt.Println("——————————————————————————————")
}

func main() {
	allStudents = make(map[uint64]*student, 48)

	for {
		// 1、打印菜单
		fmt.Println("✨ 欢迎使用函数版学生管理系统管理 ✨")
		fmt.Println(`💥 请输入序号选择你要执行的操作 💥：
👉1、查看所有学生
👉2、新增学生
👉3、删除学生
👉4、退出
	`)
		fmt.Print("👌请输入你要执行的操作,按回车确定：")
		// 2、等待用户操作
		var choice int
		fmt.Scanln(&choice)
		fmt.Printf("✅ 你选择了【 %d 】这个选项\n", choice)
		// 3、执行对应操作
		switch choice {
		case 1:
			showALlStudents()
		case 2:
			addStudent()
		case 3:
			deleteStudent()
		case 4:
			os.Exit(1) // 退出
		default:
			fmt.Println("选择错误🙅")
		}
	}
}
