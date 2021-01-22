package main

import (
	"fmt"
)

/*
如果你有50枚金币，需要分配给以下几个人：Matthew,Sarah,Augustus,Heidi,Emilie,Peter,Giana,Adriano,Aaron,Elizabeth.
分配规则如下：
a.名字中每包含1个'e',或’E‘分1枚金币
b.名字中每包含1个'i',或’I‘分2枚金币
c.名字中每包含1个'o',或’O‘分3枚金币
d.名字中每包含1个'u',或’U‘分4枚金币

写一个程序，计算每个用户分到多少金币，以及最后剩余多少金币？
程序结构如下，请实现’dispatchCoin‘函数
*/
var (
	coins = 50
	users = []string{
		"Matthew", "Sarah", "Augustus", "Heidi", "Emilie", "Peter", "Giana", "Adriano", "Aaron", "Elizabeth",
	}
	distribution = make(map[string]int, len(users))
)

func dispatchCoin() (left int) {
	//1.依次拿到每个人的名字
	for _, name := range users {
		for _, c := range name { //2.拿到每个人的名字根据规则分配金币
			switch c {
			case 'e', 'E':
				distribution[name]++ //2.1 每个人分的金币数应该保存到distribution中
				coins--
			case 'i', 'I':
				distribution[name] += 2
				coins -= 2 //2.2 记录下剩余的金币数
			case 'o', 'O':
				distribution[name] += 3
				coins -= 3
			case 'u', 'U':
				distribution[name] += 4
				coins -= 4
			}
		}
	}
	//3.整个第二部执行完就能得到最终每个人分配的金币数和剩余的金币数
	left = coins
	return
}

func main() {
	left := dispatchCoin()
	for k, v := range distribution {
		fmt.Printf("%s: 分得金币%d个\n", k, v)
	}
	fmt.Println("还剩下：", left)
}
