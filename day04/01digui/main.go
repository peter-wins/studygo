package main

import "fmt"

//递归:函数自己调用自己
//递归一定要有一个明确的退出条件

//求阶乘
// 3！= 3*2*1		==>3*2!
// 4！= 4*3*2*1		==>4*3!
// 5！= 5*4*3*2*1	==>5*4!
func f(n uint64) uint64 {
	if n <= 1 {
		return 1
	}
	return n * f(n-1)
}

//上台阶面试题
//N个台阶，一次可以走1步，也可以走2步，有多少种走法？

func step(n uint64) uint64 {
	if n == 1 { //如果只有一个台阶就只有一种走法
		return 1
	}
	if n == 2 {
		return 2
	}
	return step(n-1) + (n - 2)
}

func main() {
	//	ret := f(6)
	ret := step(6)
	fmt.Println(ret)
}
