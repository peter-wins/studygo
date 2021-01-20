package main

import "fmt"

// 数组

// 存放元素的容器
// 必须指定存放元素的类型和容量（长度）
// 数组的长度时数组类型的一部分

func main() {
	var a1 [3]bool // [true false	true]
	var a2 [4]bool // [false true false true]

	fmt.Printf("a1:%T a2:%T\n", a1, a2)

	//数组的初始化
	//如果不初始化，默认元素都是零值（布尔值：false,整型和浮点型都是0，字符串：”“）
	fmt.Println(a1, a2)

	//1.初始化方式1
	a1 = [3]bool{true, true, true}
	fmt.Println(a1)
	//2.初始化方式2 根据初始化的值自动推算数组的长度是多少
	a10 := [...]int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println(a10)
	//3.初始化方式3 根据索引来初始化
	a3 := [5]int{0: 1, 4: 2}
	fmt.Println(a3)
}
