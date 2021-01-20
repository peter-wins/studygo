package main

import "fmt"

// make()函数创造切片
func main() {
	s1 := make([]int, 5, 10) // []int:元素类型 5：长度 10：容量
	fmt.Printf("s1=%v len=%d cap=%d\n", s1, len(s1), cap(s1))

	s2 := make([]int, 0, 10)
	fmt.Printf("s2=%v len=%d cap=%d\n", s2, len(s2), cap(s2))

	//切片的赋值
	s3 := []int{1, 3, 5}
	s4 := s3 //s3和s4 都指向了同一个底层的数组
	fmt.Println(s3, s4)
	s3[0] = 1000
	fmt.Println(s3, s4)

	//切片的遍历
	//1. 索引的遍历
	for i := 0; i < len(s3); i++ {
		fmt.Println(s3[i])
	}
	//2.range遍历
	for i, v := range s3 {
		fmt.Println(i, v)
	}

}
