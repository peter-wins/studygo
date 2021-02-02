package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

/*
使用goroutine和channel实现一个计算int64随机数各位数和的程序。
开启一个goroutine循环生成int64类型的随机数，发送到jobChan
开启24个goroutine从jobChan中取出随机数计算各位数的和，将结果发送到resultChan
主goroutine从resultChan取出结果并打印到终端输出
*/

type job struct {
	value int64
}
type result struct {
	job *job
	sum int64
}

var jobChan = make(chan *job, 100)
var resultChan = make(chan *result, 100)
var wg sync.WaitGroup

func f1(ch1 chan<- *job) {
	defer wg.Done()
	//循环生成int64类型的随机数，发送到jobChan
	for {
		x := rand.Int63()
		newJob := &job{
			value: x,
		}

		ch1 <- newJob
		time.Sleep(time.Millisecond * 500)
	}
}

func f2(ch1 <-chan *job, resultChan chan<- *result) {
	defer wg.Done()
	//从jobChan中取出随机数计算各位数的和，将结果发送到resultChan
	for {
		job := <-ch1
		sum := int64(0)
		n := job.value
		for n > 0 {
			sum += n % 10
			n = n / 10
		}
		newResult := &result{
			job: job,
			sum: sum,
		}
		resultChan <- newResult
	}
}
func main() {
	wg.Add(1)
	go f1(jobChan)
	// 开启24个goroutine
	wg.Add(24)
	for i := 0; i < 24; i++ {
		go f2(jobChan, resultChan)
	}
	for result := range resultChan {
		fmt.Printf("value:%d sum:%d\n", result.job.value, result.sum)
	}
	wg.Wait()
}
