package main

import (
	"fmt"
	"math/rand"
	"time"
)

//生成一个随机数，让一个用户去猜这个数是多少？

func main() {
	var number int

	rand.Seed(time.Now().UnixNano()) //通过rand函数的Seed方法，来设置总值，这里我们以当前时间来设置总值，并且用的纳秒，十分精确了
	number = rand.Intn(100)          //随机数的范围是0-100，但不包括100
	//fmt.Println(number)
	fmt.Printf("请猜一个数字，数字的范围是：[0-100)\n")
	for {
		var input int
		fmt.Scanf("%d\n", &input) //Scanf表示让用户输入,Scanf从终端读取一个整数，并传值给input变量，&表示获取到该变量内存地址
		var flag bool = false     //通过设置flag变量，解决的是用户输入正确后可以退出
		switch {
		case number > input:
			fmt.Printf("您输入的数字太小\n")
		case number == input:
			flag = true
			fmt.Printf("恭喜您，答对了！\n")
		case number < input:
			fmt.Printf("您输入的数字太大\n")
		}
		if flag { //表示如果flag为真，则break退出这个for循环
			break
		}
	}
}
