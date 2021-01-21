package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	//1.判断字符串中汉字的数量
	s1 := "Hello沙河مرحبا안녕하십니까"
	//1.依次拿出字符串中的汉字
	var count int
	for _, c := range s1 {
		//2.判断当前这个字符是不是汉字
		if unicode.Is(unicode.Han, c) {
			count++
		}
	}
	//3.把汉字出现的次数累加
	fmt.Println(count)

	//2.写一个程序，统计一个字符串中每个单词出现的次数。比如：”how do you do”中how=1 do=2 you=1。
	s2 := "how do you do"
	//2.1把字符串按照空格切割得到切片
	s3 := strings.Split(s2, " ")
	//fmt.Println(s2, s3)
	//2.2遍历切片存储到一个map
	m1 := make(map[string]int, 10)
	for _, w := range s3 {
		//fmt.Println(w)
		//2.3如果原来的map中不存在w这个key那么出现次数=1
		if _, ok := m1[w]; !ok {
			m1[w] = 1
		} else {
			m1[w]++ //2.4如果map中存在w这个key，那么出现次数+1
		}
	}
	//累加出现的次数
	for k, v := range m1 {
		fmt.Println(k, v)
	}

	//回文判断
	ss := "上海自来水来自海上"
	r := make([]rune, 0, len(ss))
	for _, c := range ss {
		r = append(r, c)
	}
	//fmt.Println(r)
	for i := 0; i < len(r)/2; i++ {
		if r[i] != r[len(r)-1-i] {
			fmt.Println("不是回文")
			return
		}
	}
	fmt.Println("是回文")
}
