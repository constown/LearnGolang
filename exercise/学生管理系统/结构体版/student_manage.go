package main

import "fmt"

// 学生管理系统
// 有一个物件
// 1.他保存了一些数据   --> 结构体的字段
// 2.他有三个功能  	 --> 结构体的方法

type student struct {
	id   int64
	name string
}

// 造一个学生的管理对象
type studentManage struct {
	allStudent map[int64]student
}

// 查看学生
func (s studentManage) showStudent() {
	// s.allStudent 这个map中吧所有学生逐个打印
	for _, value := range s.allStudent {
		fmt.Printf("学号: %d 姓名:%s\n", value.id, value.name)
	}
}

// 添加学生
func (s studentManage) addStudent() {
	// 根据用户的输入创建新的学生
	var (
		id   int64
		name string
	)
	// 获取用户输入
	fmt.Print("请输入姓名:")
	fmt.Scanln(&name)
	id = int64(len(s.allStudent))
	// 根据用户输入去创建结构体对象
	newStudent := student{
		id:   id,
		name: name,
	}
	// 把新的学生放进map中
	s.allStudent[id] = newStudent
}

// 删除学生
func (s studentManage) deleteStudent() {
	fmt.Print("请输入要删除的学生学号:")
	var (
		deleteId int64
	)
	fmt.Scanln(&deleteId)
	value, ok := s.allStudent[deleteId]
	if !ok {
		fmt.Println("查无此人")
		return
	}
	delete(s.allStudent, deleteId)
	fmt.Printf("删除的学生的信息如下:学号:%d, 名字: %s\n", value.id, value.name)
}

// 修改学生
func (s studentManage) editStudent() {
	// 获取用户输入的学号
	fmt.Print("请输入要修改的学生学号:")
	var editId int64
	fmt.Scanln(&editId)
	// 展示该学号的信息,如果没有就提示查无此人
	value, ok := s.allStudent[editId] // ok 等于 true 表示有值 false 代表没有值  取不到值的时候 value 是 对应类型的 零值
	if !ok {
		fmt.Println("查无此人")
		return
	}
	fmt.Printf("你要修改的学生的信息如下:学号:%d, 名字: %s\n", value.id, value.name)
	fmt.Print("请输入新的名字:")
	var newName string
	fmt.Scanln(&newName)
	value.name = newName
	s.allStudent[editId] = value
}
