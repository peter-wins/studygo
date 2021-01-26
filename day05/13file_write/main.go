package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

//打开文件写内容
func writeDemo1() {
	fileObj, err := os.OpenFile("test.txt", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		fmt.Printf("open file failed, err:%v\n", err)
		return
	}
	//write
	fileObj.Write([]byte("这是一个测试内容！\n"))
	//writestring
	fileObj.WriteString("又是一个测试内容！")
	fileObj.Close()
}
func writeDemo2() {
	fileObj, err := os.OpenFile("test.txt", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Printf("open file failed, err:%v\n", err)
		return
	}
	defer fileObj.Close()
	//创建一个写的对象
	wr := bufio.NewWriter(fileObj)
	wr.WriteString("Hello 北京！\n") //写到缓存中
	wr.Flush()                    // 将缓存中的内容写入文件
}

func writeDemo3() {
	str := "Hello 沙河！！"
	err := ioutil.WriteFile("./test.txt", []byte(str), 0666)
	if err != nil {
		fmt.Printf("open file failed, err:%v\n", err)
		return
	}
}
func main() {
	//writeDemo1()
	//writeDemo2()
	writeDemo3()
}
