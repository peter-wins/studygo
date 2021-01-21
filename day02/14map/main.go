package main

import "fmt"

// map

func main() {
	var m1 map[string]int
	fmt.Println(m1 == nil) //没有初始化的map
	m1 = make(map[string]int, 10)
	m1["wins"] = 20
	m1["jiwuming"] = 30
	fmt.Println(m1)
	fmt.Println(m1["jiwuming"])

	//约定俗成用OK接受返回的布尔值
	fmt.Println(m1["娜扎"])
	v, ok := m1["娜扎"]
	if !ok {
		fmt.Println("查无此key")
	} else {
		fmt.Println(v)
	}

	//map遍历
	for k, v := range m1 {
		fmt.Println(k, v)
	}
	//只遍历KEY
	for k := range m1 {
		fmt.Println(k)
	}
	//只遍历value
	for _, v := range m1 {
		fmt.Println(v)
	}

	//删除
	delete(m1, "jiwuming")
	fmt.Println(m1)
}
