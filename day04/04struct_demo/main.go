package main

import "fmt"

//结构体是值类型
type person struct {
	name, gender string
}

//Go语言中函数参数永远是copy
func f(x person) {
	x.gender = "女" //修改的是副本的gender
}

func main() {
	var p person
	p.name = "周林"
	p.gender = "男"
	f(p)
	fmt.Println(p.gender)
}
