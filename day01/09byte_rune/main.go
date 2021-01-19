package main

import "fmt"

func main() {
	s := "Hello 沙河"
	//	len()求byte字节的数量
	n := len(s) //	求字符串s的长度，把长度保存再变量n中
	fmt.Println(n)

	// for i := 0; i < len(s); i++ {
	// 	fmt.Println(s[i])
	// 	fmt.Printf("%c\n", s[i]) // %c字符
	// }

	// for _, c := range s { //从字符串中拿出具体的字符
	// 	fmt.Printf("%c\n", c)
	// }

	//	字符串修改
	s2 := "白萝卜"      // ==> ‘白’ ‘萝’ ‘卜’
	s3 := []rune(s2) //	把字符串强制转换成一个rune切片
	s3[0] = '红'
	fmt.Println(string(s3)) //	rune切片强制转换成字符串

	c1 := "红"
	c2 := '红'
	fmt.Printf("c1:%T c2:%T\n", c1, c2)
	c3 := "H"       // string
	c4 := byte('H') //byte(uint8)
	fmt.Printf("c3:%T c4:%T\n", c3, c4)

	//	类型转换
	n1 := 10 // int
	var f float64
	f = float64(n1)
	fmt.Println(f)
	fmt.Printf("%T\n", f)
}
