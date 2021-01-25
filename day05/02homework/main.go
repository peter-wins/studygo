package main

import (
	"fmt"
	"os"
)

//学生管理系统
//有一个物件
// 1.它保存了一些数据	————>结构体的字段
// 2.他有4个功能	————>结构体的方法

var smr studentMgr

func showMenu() {
	fmt.Println("-------welcome to sms-------")
	fmt.Println(`
	1.查看所有学生
	2.添加学生
	3.修改学生
	4.删除学生
	5.退出系统
	`)
}

// 造一个学生的管理者
func main() {
	smr = studentMgr{
		allStudent: make(map[int64]student, 100),
	}
	for {
		showMenu()
		fmt.Print("请输入序号：")
		var choice int
		fmt.Scanln(&choice)
		fmt.Println("你的输入是：", choice)

		switch choice {
		case 1:
			smr.showStudents()
		case 2:
			smr.addStudent()
		case 3:
			smr.editStudent()
		case 4:
			smr.deleteStudent()
		case 5:
			os.Exit(1)
		default:
			fmt.Println("滚犊子~~~")
		}
	}
}
