package main

import "fmt"

//defer
//defer多用于函数结束之前释放资源(文件句柄、数据库连接、socket链接)
func deferDemo() {
	fmt.Println("start")
	fmt.Println("嘿嘿嘿") //defer把后面的语句延迟到函数即将返回的时候在执行
	fmt.Println("哈哈哈") //一个函数中可以有多个defer2语句
	fmt.Println("呵呵呵") //多个defer语句按照先进后出(后进先出)的顺序延迟执行
	fmt.Println("end")
}
func main() {
	deferDemo()
}
