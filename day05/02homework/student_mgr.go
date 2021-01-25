package main

import "fmt"

//学生管理系统
//有一个物件
// 1.它保存了一些数据	————>结构体的字段
// 2.他有4个功能	————>结构体的方法

type student struct {
	id   int64
	name string
}

// 造一个学生的管理者
type studentMgr struct {
	allStudent map[int64]student
}

//查看学生
func (s studentMgr) showStudents() {
	for _, stu := range s.allStudent {
		fmt.Printf("学号：%d 姓名：%s\n", stu.id, stu.name)
	}
}

// 增加学生
func (s studentMgr) addStudent() {
	var (
		stuID   int64
		stuName string
	)
	fmt.Print("请输入学号：")
	fmt.Scanln(&stuID)
	fmt.Print("请输入姓名：")
	fmt.Scanln(&stuName)
	//根据用户输入创建结构体对象
	newStu := student{
		id:   stuID,
		name: stuName,
	}
	//把新的学生放到s.allstudent这个map中
	s.allStudent[stuID] = newStu
	fmt.Println("添加成功!")
}

// 修改学生
func (s studentMgr) editStudent() {
	var stuID int64
	fmt.Print("请输入学号：")
	fmt.Scanln(&stuID)
	//展示该学号对应的学生信息，如果没有此人提示查无此人
	stuObj, ok := s.allStudent[stuID]
	if !ok {
		fmt.Println("查无此人")
		return
	}
	fmt.Printf("你要修改的学生信息如下：学号：%d 姓名：%s\n", stuObj.id, stuObj.name)
	fmt.Print("请输入学生的新名字：")
	var newName string
	fmt.Scanln(&newName)
	stuObj.name = newName
	//更新学生的姓名
	s.allStudent[stuID] = stuObj
	fmt.Println("更新成功！")
}

// 删除学生
func (s studentMgr) deleteStudent() {
	var stuID int64
	fmt.Print("请输入要删除的学生学号：")
	fmt.Scanln(&stuID)
	_, ok := s.allStudent[stuID]
	if !ok {
		fmt.Println("查无此人")
		return
	}
	//从map中删除键值对
	delete(s.allStudent, stuID)
	fmt.Println("删除成功！")

}
