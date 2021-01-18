package main

import "fmt"

// Go语言推荐使用驼峰式命名

// var student_name string
// var studentName string
// var StudentName string

//声明变量
//var name string
//var age int
//var isOk bool

//批量声明
var (
	name string
	age  int
	isOk bool
)

func main() {
	name = "理想"
	age = 16
	isOk = true
	//Go语言中声明的非全局变量必须使用，不使用编译不过去

	fmt.Print(isOk)               // 在终端输出这段内容
	fmt.Println()                 //换行
	fmt.Printf("name:%s\n", name) // %s占位符 使用name这个变量的值去替换占位符 \n换行
	fmt.Println(age)              // 打印完指定的内容之后会在后面加一个换行符

	// 声明变量同时赋值
	var s1 string = "wins"
	fmt.Println(s1)
	// 类型推导（根据值判断变量是什么类型）
	var s2 = "20"
	fmt.Println(s2)
	// 简短变量声明，只能在函数里面用
	s3 := "heiheihei"
	fmt.Println(s3)
}

//匿名变量
//在使用多重赋值时，如果想要忽略某个值，可以使用匿名变量（anonymous variable）。 匿名变量用一个下划线_表示，例如：

// func foo() (int, string) {
// 	return 10, "Q1mi"
// }
// func main() {
// 	x, _ := foo()
// 	_, y := foo()
// 	fmt.Println("x=", x)
// 	fmt.Println("y=", y)
// }

/*
	匿名变量不占用命名空间，不会分配内存，所以匿名变量之间不存在重复声明。 (在Lua等编程语言里，匿名变量也被叫做哑元变量。)

	注意事项：

	函数外的每个语句都必须以关键字开始（var、const、func等）
	:=不能使用在函数外。
	_多用于占位，表示忽略值。
*/
