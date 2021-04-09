package main

import (
	"fmt"
	"sync"
	"time"
)

// 为什么需要context
var wg sync.WaitGroup
var exitChan = make(chan bool, 1)

func f() {
	defer wg.Done()
LOOP:
	for {
		fmt.Println("熊大")
		time.Sleep(time.Millisecond * 500)
		select { // 多路复用
		case <-exitChan:
			break LOOP
		default:
		}
	}
}

func main() {
	wg.Add(1)
	go f()
	time.Sleep(time.Second * 5)
	// 如何通知子groutine退出
	exitChan <- true
	wg.Wait()
}
