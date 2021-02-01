package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func a() {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		fmt.Printf("A:%d\n", i)
	}
}
func b() {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		fmt.Printf("B:%d\n", i)
	}
}
func main() {
	runtime.GOMAXPROCS(1)
	fmt.Printf("我机器的有%d个逻辑处理器\n", runtime.NumCPU())
	wg.Add(2)
	go a()
	go b()
	wg.Wait()
}
