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

//break(跳出循环)
//break语句可以结束for、switch和select的代码块。

//break语句还可以在语句后面添加标签，表示退出某个标签对应的代码块，标签要求必须定义在对应的for、switch和 select的代码块上。 举个例子：

func breakDemo1() {
BREAKDEMO1:
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if j == 2 {
				break BREAKDEMO1
			}
			fmt.Printf("%v-%v\n", i, j)
		}
	}
	fmt.Println("...")
}

//continue(继续下次循环)
//continue语句可以结束当前循环，开始下一次的循环迭代过程，仅限在for循环内使用。

//在 continue语句后添加标签时，表示开始标签对应的循环。例如：

func continueDemo() {
forloop1:
	for i := 0; i < 5; i++ {
		// forloop2:
		for j := 0; j < 5; j++ {
			if i == 2 && j == 2 {
				continue forloop1
			}
			fmt.Printf("%v-%v\n", i, j)
		}
	}
}
