package main

import "fmt"

// array练习题
// 求数组[1,3,5,7,8]所有元素的和
func main() {
	a1 := [...]int{1, 3, 5, 7, 8}
	// 把数组中的每个元素遍历出来，累加到一个sum变量就可以了
	sum := 0
	for _, v := range a1 {
		sum = sum + v
	}
	fmt.Println(sum)

	// 找出数组中和为指定值的两个元素的下标，比如从数组[1,3,5,7,8]中找出和为8的两个元素的下标分别为(0,3)和(1,2)
	// 定义两个for循环，外层的从第一个元素开始遍历
	// 内层的for循环从外层后面的哪个开始
	// 他们俩个数的和为8
	for i := 0; i < len(a1); i++ {
		for j := i + 1; j < len(a1); j++ {
			if a1[i]+a1[j] == 8 {
				fmt.Printf("(%d,%d)\n", i, j)
			}
		}
	}
}
