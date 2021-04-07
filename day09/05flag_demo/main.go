package main

import (
	"flag"
	"fmt"
)

// flag 获取命令行参数

func main() {
	// 创建一个标志位参数
	name := flag.String("name", "小明", "姓名")
	age := flag.Int("age", 80, "请输入真实年龄")
	// 使用
	flag.Parse()
	fmt.Println(*name)
	fmt.Println(*age)
}
