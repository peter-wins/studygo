package main

import "fmt"

//引出接口的实例

//定义一个能叫的类型
type speaker interface {
	speak() //只要实现了speck的方法的变量都是specker类型	方法签名 可以有多个
}
type cat struct {
}
type dog struct {
}

func (c cat) speak() {
	fmt.Println("喵喵喵~~")

}
func (d dog) speak() {
	fmt.Println("汪汪汪~~")
}
func da(x speaker) {
	//接收一个参数，传进来什么就打什么
	x.speak()
}

func main() {
	var c1 cat
	var d1 dog

	da(c1)
	da(d1)
}
