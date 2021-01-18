package main

import "fmt"

//	字符串，Go语言中字符串使用双引号包裹的
//	Go语言中单引号包裹的是字符

//	s := "Hello world"
//	// 单独的字母、汉字、符号表示一个字符
//	c1 := 'h'
//	c1 := '1'
//	c1 := '撒'
//	字节：1个字节=8bit（8个二进制位）
//	1个字符‘A’=1个字节
//	1个utf8编码的汉字 ‘啥’= 一般占3个字节
func main() {
	//	\ 本来是具有特殊含义的，我应该该告诉程序我写的就是一个单纯的 \
	path := "\"D:\\work\\go\\src\\github.com\\peter-wins\\studygo\\day01\""
	fmt.Println(path)

	s := "I’m ok"
	fmt.Println(s)

	// 多行的字符串

	s2 := `
		世情薄
		人情恶
		雨送黄昏花易落
	`
	fmt.Println(s2)
	s3 := `D:\work\go\src\github.com\peter-wins\studygo\day01`
	fmt.Println(s3)

	// 字符串相关操作

	fmt.Println(len(s3)) // 求字符串长度

	// 字符串拼接

	name := "理想"
	world := "大傻逼"

	ss := name + world
	fmt.Println(ss)
	ss1 := fmt.Sprintf("%s%s", name, world)
	fmt.Println(ss1)
}
