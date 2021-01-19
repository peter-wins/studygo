package main

import "fmt"

//goto

func main() {
	// var flag = false
	// for i := 0; i < 10; i++ {
	// 	for j := 'A'; j < 'Z'; j++ {
	// 		if j == 'C' {
	// 			flag = true
	// 			break
	// 		}
	// 		fmt.Printf("%v-%c\n", i, j)
	// 	}
	// 	if flag {
	// 		break
	// 	}
	// }

	//goto+label实现跳出多层for循环
	for i := 0; i < 10; i++ {
		for j := 'A'; j < 'Z'; j++ {
			if j == 'C' {
				goto XX //跳到我指定的标签
			}
			fmt.Printf("%v-%c\n", i, j)
		}
	}
XX: //label标签
	fmt.Println("over")
}
