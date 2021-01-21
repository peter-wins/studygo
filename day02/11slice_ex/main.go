package main

import (
	"fmt"
	"sort"
)

//1.请写出下面代码的输出结果。

func main() {
	var a = make([]int, 5, 10)
	fmt.Println(a)
	for i := 0; i < 10; i++ {
		a = append(a, i)
	}
	fmt.Println(a)

	//2.请使用内置的sort包对数组var a = [...]int{3, 7, 8, 9, 1}进行排序
	var a1 = [...]int{3, 7, 8, 9, 1}
	sort.Ints(a1[:])
	fmt.Println(a1)

}
