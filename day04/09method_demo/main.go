package main

import "fmt"

//自定义类型加方法
//不能给别的包里的类型添加方法，只能给自己的包里的类型添加方法
type myInt int

func (m myInt) hello() {
	fmt.Println("我是一个int")
}
func main() {
	m := myInt(100)
	m.hello()
}
