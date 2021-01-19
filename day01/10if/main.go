package main

import "fmt"

//	if条件判断

func main() {
	a1 := 16
	if a1 > 18 {
		fmt.Println("澳门首家赌场开业啦！")
	} else {
		fmt.Println("回家写作业去！")
	}

	//	多个条件判断
	a2 := 36
	if a2 > 35 {
		fmt.Println("人到中年")
	} else if a2 > 18 {
		fmt.Println("还是青年")
	} else {
		fmt.Println("好好学习")
	}

	//	作用域
	//	a3变量此时只在if条件判断语句中生效
	if a3 := 20; a3 > 18 {
		fmt.Println("澳门首家赌场开业啦！")
	} else {
		fmt.Println("回家写作业去！")
	}
}
