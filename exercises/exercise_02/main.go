package main

import (
	"fmt"
	"math/rand"
	"time"
)

//求数组所有元素之和
//方法一
func f1() {
	a := [...]int{1, 3, 5, 10}
	sum := 0
	for i := 0; i < len(a); i++ {
		sum += a[i]
	}
	fmt.Println(sum)

}

//方法二
func f2() {
	var arr [10]int
	rand.Seed(time.Now().UnixNano()) //通过rand函数的Seed方法，来设定随机数总值
	for i := 0; i < 10; i++ {
		arr[i] = rand.Intn(1000) //数组取值为随机数，随机数范围为[1-10000)
	}

	var sum int
	for i := 0; i < len(arr); i++ {
		sum += arr[i]
	}
	fmt.Println("sum =", sum)
}

func main() {
	f1()
	f2()

}
