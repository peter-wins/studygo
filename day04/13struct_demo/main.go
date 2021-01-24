package main

import "fmt"

//结构体嵌套
type address struct {
	province string
	city     string
}
type workplace struct {
	province string
	city     string
}
type person struct {
	name string
	age  int
	//	workplace //匿名嵌套结构体
	address //匿名嵌套结构体
	//addr address
}
type company struct {
	name    string
	address //匿名嵌套结构体
	//addr address
}

func main() {
	p1 := person{
		name: "周玲",
		age:  19,
		address: address{
			province: "青海",
			city:     "西宁",
		},
	}
	fmt.Println(p1)
	fmt.Printf("%s:今年%d岁,来自 %s.%s\n", p1.name, p1.age, p1.address.province, p1.address.city)
	fmt.Println(p1.city) //先在自己结构体找这个字段，找不到就去匿名嵌套的结构体中查找该字段
}
