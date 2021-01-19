package main

import "fmt"

func main() {
	//	基本格式
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	//	变种1
	var j = 5
	for ; j < 10; j++ {
		fmt.Println(j)
	}

	//	变种2
	var k = 6
	for k < 10 {
		fmt.Println(k)
		k++
	}

	//	无限循环

	// for {
	// 	fmt.Println("123")
	// }

	//	for range 循环
	s := "Hello沙河"
	for i, v := range s {
		fmt.Printf("%d %c\n", i, v)
	}
}
