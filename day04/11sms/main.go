package main

import (
	"fmt"
	"os"
)

/*
	函数版学生管理系统
	写一个系统能查看、增加、删除学生
*/

type student struct {
	id   int64
	name string
}

var allStudent = make(map[int64]*student, 50) //声明变量

//newStudent 是student类型的构造函数
func newStudent(id int64, name string) *student {
	return &student{
		id:   id,
		name: name,
	}
}
func showAllstudent() {
	//遍历
	//fmt.Println("我要打印这个map！")
	//fmt.Println(allStudent)
	for k, v := range allStudent {
		fmt.Printf("学号：%d 姓名：%s\n", k, v.name)
	}
}
func addStudent() {
	//向allstudent中添加一个新的学生
	//1.创建一个新学生
	var (
		id   int64
		name string
	)
	//1.1获取用户输入
	fmt.Print("请输入学生学号：")
	fmt.Scanln(&id)
	fmt.Print("请输入学生姓名：")
	fmt.Scanln(&name)
	//1.2造学生(调用student的构造函数)
	newStu := newStudent(id, name)
	//2.追加到allStudent这个map中
	allStudent[id] = newStu
}
func deleteStudent() {
	//输入要删除学生的学号
	var deleteID int64
	fmt.Print("请输入要删除的学生学号：")
	fmt.Scanln(&deleteID)
	//在map中删除学号对应的键值对
	delete(allStudent, deleteID)
}

func main() {
	//allStudent = make(map[int64]*student, 48) //初始化map
	for {
		//1.打印菜单
		fmt.Println("=====欢迎使用学生管理系统======")
		fmt.Println(`
	1.查看所有学生
	2.新增学生
	3.删除学生
	4.退出系统
		`)
		fmt.Print("请输入选项：")
		//2.等待用户输入
		var choise int
		fmt.Scanln(&choise)
		fmt.Printf("你选择了%d这个选项\n", choise)
		//3.执行对应的函数
		switch choise {
		case 1:
			showAllstudent()
		case 2:
			addStudent()
		case 3:
			deleteStudent()
		case 4:
			os.Exit(1)
		default:
			fmt.Println("滚犊子~")
		}
	}
}
