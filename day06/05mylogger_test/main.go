package main

import (
	"time"

	"github.com/peter-wins/studygo/day06/mylogger"
)

// 测试我们自己写的日志库
func main() {
	log := mylogger.NewLog("info")
	for {
		log.Debug("这是一条Debug级别的日志~~~")
		log.Info("这是一条Info级别的日志~~~")
		log.Warning("这是一条Warning级别的日志~~~")
		log.Error("这是一条Error级别的日志~~~")
		log.Fatal("这是一条Fatal级别的日志~~~")
		time.Sleep(time.Second)
	}
}
