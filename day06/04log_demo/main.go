package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

// log demo

func main() {
	fileObj, err := os.OpenFile("./test.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("open file failed, err:%v", err)
		return
	}
	log.SetOutput(fileObj)
	for {
		log.Println("This is a test log...")
		time.Sleep(time.Second * 2)
	}
}
