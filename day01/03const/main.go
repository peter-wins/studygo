package main

import "fmt"

//常量
// 定义了常量之后不能修改
// 在程序运行期间不会改变的量
const pi = 3.141596

// 批量声明常量
const (
	statusOK = 200
	notFound = 404
)

// const同时声明多个常量时，如果省略了值则表示和上面一行的值相同
const (
	n1 = 100
	n2
	n3
)

//iota 在const关键字出现时将被重置为0。const中每新增一行常量声明将使iota计数一次(iota可理解为const语句块中的行索引)。
//使用iota能简化定义，在定义枚举时很有用
const (
	a1 = iota //0
	a2        //1
	a3        //2
)

//几个常见的iota示例:
//使用_跳过某些值

const (
	b1 = iota //0
	b2        //1
	_
	b4 //3
)

//iota声明中间插队

const (
	c1 = iota //0
	c2 = 100  //100
	c3 = iota //2
	c4        //3
)
const c5 = iota //0

//多个iota定义在一行
const (
	d1, d2 = iota + 1, iota + 2 //d1:1	d2:2
	d3, d4 = iota + 1, iota + 2 //d3:2	d4:3
)

//定义数量级
const (
	_  = iota
	KB = 1 << (10 * iota)
	MB = 1 << (10 * iota)
	GB = 1 << (10 * iota)
	TB = 1 << (10 * iota)
	PB = 1 << (10 * iota)
)

func main() {
	fmt.Println("n1:", n1)
	fmt.Println("n2:", n2)
	fmt.Println("n3:", n3)

	fmt.Println("a1:", a1)
	fmt.Println("a2:", a2)
	fmt.Println("a3:", a3)

	fmt.Println("d1=", d1)
	fmt.Println("d2=", d2)
	fmt.Println("d3=", d3)
}
