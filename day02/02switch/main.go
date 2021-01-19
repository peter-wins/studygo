package main

import "fmt"

//switch 简化大量的判断(一个变量和具体的值作比较)

func main() {
	//	var n = 4
	switch n := 5; n {
	case 1:
		fmt.Println("大拇指")
	case 2:
		fmt.Println("食指")
	case 3:
		fmt.Println("中指")
	case 4:
		fmt.Println("无名指")
	case 5:
		fmt.Println("小拇指")
	default:
		fmt.Println("无效的输入")
	}

	// 一个分支可以有多个值，多个case值中间使用英文逗号分隔
	switch n := 6; n {
	case 1, 3, 5, 7, 9:
		fmt.Println("奇数")
	case 0, 2, 4, 6, 8:
		fmt.Println("偶数")
	}

	//分支还可以使用表达式，这时候switch语句后面不需要再跟判断变量。例如：
	age := 10
	switch {
	case age < 25:
		fmt.Println("好好学习吧")
	case age > 25 && age < 35:
		fmt.Println("好好工作吧")
	case age > 60:
		fmt.Println("好好享受吧")
	default:
		fmt.Println("活着真好")
	}
}
