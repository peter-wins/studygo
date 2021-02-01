package main

import (
	"fmt"
	"strconv"
)

func main() {
	// 从字符串中解析出对应的数据
	str := "10000"
	//ret1 := int64(str)
	ret1, err := strconv.ParseInt(str, 10, 64) // 10表示十进制，64表示int64
	if err != nil {
		fmt.Printf("parseint failed, err:%v\n", err)
		return
	}
	fmt.Printf("%#v %T\n", ret1, ret1)
	// Atoi 字符串转换成int
	retInt, _ := strconv.Atoi(str)
	fmt.Println(retInt)

	// 从字符串中解析出布尔值
	boolStr := "true"
	boolValue, _ := strconv.ParseBool(boolStr)
	fmt.Printf("%#v %T\n", boolValue, boolValue)

	// 把数字转换成字符串类型

	//i := int32(2316)
	i := 100
	//ret2 := string(i)
	ret2 := fmt.Sprintf("%d", i)
	fmt.Printf("%#v\n", ret2)

	ret3 := strconv.Itoa(i)
	fmt.Printf("%#v\n", ret3)
}
