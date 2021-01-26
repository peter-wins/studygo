package main

import "fmt"

func main() {
	var a interface{}
	a = 100
	fmt.Printf("%T\n", a)
	//r如何判断a保存的值是什么类型的
	//类型断言
	//1.x(T)
	v1, ok := a.(int8)
	if ok {
		fmt.Println("猜对了，a是int8", v1)
	} else {
		fmt.Println("猜错了，a不是int8")
	}

	//switch
	switch v2 := a.(type) {
	case int8:
		fmt.Println("int8", v2)
	case int16:
		fmt.Println("int16", v2)
	case string:
		fmt.Println("string", v2)
	case int:
		fmt.Println("int", v2)
	default:
		fmt.Println("滚~~~")
	}
}
