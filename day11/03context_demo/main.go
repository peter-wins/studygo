package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// 为什么需要context
var wg sync.WaitGroup
var exitChan = make(chan bool, 1)

func f2(ctx context.Context) {
	defer wg.Done()
LOOP:
	for {
		fmt.Println("熊二")
		time.Sleep(time.Millisecond * 500)
		select { // 多路复用
		case <-ctx.Done():
			break LOOP
		default:
		}
	}
}

func f(ctx context.Context) {
	defer wg.Done()
	go f2(ctx)
LOOP:
	for {
		fmt.Println("熊大")
		time.Sleep(time.Millisecond * 500)
		select { // 多路复用
		case <-ctx.Done():
			break LOOP
		default:
		}
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	wg.Add(1)
	go f(ctx)
	time.Sleep(time.Second * 5)
	// 如何通知子groutine退出
	cancel()
	wg.Wait()
}
