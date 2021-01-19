package main

import "fmt"

// 运算符
func main() {
	var (
		a = 5
		b = 2
	)
	//算术运算符
	fmt.Println(a + b)
	fmt.Println(a - b)
	fmt.Println(a * b)
	fmt.Println(a / b)
	fmt.Println(a % b)
	a++ // 单独的语句，不能放在=的右边赋值 ==> a=a+1
	b-- // 单独的语句，不能放在=的右边赋值 ==> b=b-1
	fmt.Printf("a++的值为%d\n", a)
	fmt.Printf("b--的值为%d\n", b)

	//关系运算符
	fmt.Println(a == b) //Go语言是强类型，相同类型的变量才能比较
	fmt.Println(a != b)
	fmt.Println(a >= b)
	fmt.Println(a > b)
	fmt.Println(a <= b)
	fmt.Println(a < b)
}
