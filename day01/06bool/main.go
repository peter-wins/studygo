package main

import "fmt"

// 布尔值

func main() {
	b1 := true
	var b2 bool // 默认是false
	fmt.Printf("%T\n", b1)
	fmt.Printf("类型是:%T value:%v\n", b2, b2) // v%打印变量的值 可以是任意的类型
}
