package main

import "fmt"

//结构体
type person struct {
	name   string
	age    int
	gender string
	hobby  []string
}

func main() {
	var p person //声明一个person类型的变量P
	//通过字段赋值
	p.name = "周琳"
	p.age = 30
	p.gender = "女"
	p.hobby = []string{"画画", "做饭", "拖地"}
	fmt.Println(p)
	fmt.Printf("%T\n", p)
	fmt.Println(p.name)

	//匿名结构体：多用于临时场景

	var s struct {
		x string
		y int
	}
	s.x = "哈哈"
	s.y = 100

	fmt.Printf("type:%T value:%v\n", s, s)
}
