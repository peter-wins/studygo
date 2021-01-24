package main

import (
	"encoding/json"
	"fmt"
)

//结构体与json
//序列化：把Go语言中的结构体与变量-->json格式的字符串
//反序列化：json格式的字符串 --> Go语言中能识别的结构体变量

type person struct {
	Name string `json:"name" db:"name" ini:"name"`
	Age  int    `json:"age"`
}

func main() {
	p1 := person{
		Name: "周玲",
		Age:  18,
	}
	//序列化
	b, err := json.Marshal(p1)
	if err != nil {
		fmt.Printf("marshal failed,err:%v", err)
		return
	}
	fmt.Printf("%v\n", string(b))

	//反序列化
	str := `{"name":"周玲","age":18}`
	var p2 person
	json.Unmarshal([]byte(str), &p2) //穿指针是为了能在json.unmarshal内部修改p2的值
	fmt.Printf("%v\n", p2)
}
