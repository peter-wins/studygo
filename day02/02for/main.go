package main

import "fmt"

//	流程控制之跳出for循环

func main() {
	for i := 0; i < 10; i++ {
		if i == 5 { //	当i=5时跳出for循环
			break //	跳出for循环
		}
		fmt.Println(i)
	}
	fmt.Println("over")
}
