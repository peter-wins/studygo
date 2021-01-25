package main

import "fmt"

// 接口实例2
//不管是什么牌子的车，都能跑
type car interface {
	run()
}
type ferrari struct {
	brand string
}
type porsche struct {
	brand string
}

//造方法
func (f ferrari) run() {
	fmt.Printf("%s的速度70迈~~\n", f.brand)
}
func (p porsche) run() {
	fmt.Printf("%s的速度80迈~~\n", p.brand)
}

//drive函数接收一个car类型的变量
func drive(c car) {
	c.run()
}

func main() {
	var f1 = ferrari{
		brand: "法拉利",
	}
	var p1 = porsche{
		brand: "保时捷",
	}
	drive(f1)
	drive(p1)
}
