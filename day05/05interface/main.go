package main

import "fmt"

//接口的实现
type animal interface {
	move()
	eat(string)
}
type cat struct {
	name string
	feet int8
}

func (c cat) move() {
	fmt.Println("走猫步...")
}

func (c cat) eat(food string) {
	fmt.Printf("猫吃%s...\n", food)
}

type chicken struct {
	name string
	feet int8
}

func (c chicken) move() {
	fmt.Println("鸡动...")
}

func (c chicken) eat(food string) {
	fmt.Printf("鸡吃%s...\n", food)
}

func main() {
	var a1 animal //	定义一个接口类型的变量
	bc := cat{
		name: "淘气",
		feet: 4,
	}

	a1 = bc
	a1.eat("小黄鱼")
	fmt.Println(a1)
}
