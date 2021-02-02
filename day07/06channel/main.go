package main

import (
	"fmt"
	"sync"
)

var a []int
var b chan int // 需要指定通道中元素的类型
var wg sync.WaitGroup

func nobufChannel() {
	fmt.Println(b)     //nil
	b = make(chan int) // 初始化通道 不带缓冲区的
	wg.Add(1)
	go func() {
		defer wg.Done()
		x := <-b
		fmt.Println("后台goroutine从通道b中取到了", x)
	}()
	b <- 10 //deadlock
	fmt.Println("10发送到通道b中了...")
	wg.Wait()
}

func bufChannel() {
	fmt.Println(b)        //nil
	b = make(chan int, 1) // 初始化通道 带缓冲区的
	b <- 10
	fmt.Println("10发送到通道b中了...")
}
func main() {
	bufChannel()
}
