package main

import "fmt"

func f1() {
	fmt.Println("Hello沙河")
}
func f2(name string) {
	fmt.Println("Hello", name)
}
func f3(x int, y int) int {
	sum := x + y
	return sum
}
func f4(x, y int) int {
	return x + y
}

//可变参数
func f5(title string, y ...int) (sum int) {
	fmt.Println(y) //y是一个int类型的切片
	return 1
}

// 命名返回值
func f6(x, y int) (sum int) {
	sum = x + y //如果使用命名的返回值，那么函数中可以直接使用返回值变量
	return      //如果使用命名的返回值，return后面可以省略返回值变量
}

//Go语言中支持多个返回值
func f7(x, y int) (sum, sub int) {
	sum = x + y
	sub = x - y
	return
}
func main() {
	f1()
	f2("world")
	fmt.Println(f3(100, 200)) //调用函数
	ret := f3(100, 200)
	fmt.Println(ret)
	f4(1, 2)
	f5("lixiang", 1, 2, 3, 4, 5)
	f6(1, 2)
	f7(30, 50)
}
