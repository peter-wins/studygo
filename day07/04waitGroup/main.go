package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func f() {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 10; i++ {
		r1 := rand.Int()    // int64
		r2 := rand.Intn(10) // 0 <= x < 10
		fmt.Println(r1, r2)
	}
}

var wg sync.WaitGroup

func f1(i int) {
	defer wg.Done()
	time.Sleep(time.Second * time.Duration(rand.Intn(3)))
	fmt.Println(i)
}

func main() {
	// f()
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go f1(i)
	}
	wg.Wait() // 等待wg的计数器减为0
}
