package main

import "fmt"

//pointer
func main() {
	//1. &：取内存地址
	n := 18
	p := &n
	fmt.Println(p)
	fmt.Printf("%T\n", p) //*int: int类型的指针
	//2. *：根据地址取值
	m := *p
	fmt.Println(m)
	fmt.Printf("%T\n", m)

	var a1 *int //nil pionter
	fmt.Println(a1)
	var a2 = new(int) //申请一个内存地址
	fmt.Println(a2)
	fmt.Println(*a2)
	*a2 = 100
	fmt.Println(*a2)
}
