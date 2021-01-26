package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

//打开文件
func readFromFile1() {
	fileObj, err := os.Open("./main.go")
	if err != nil {
		fmt.Printf("open file failed, err:%v\n", err)
		return
	}

	//关闭文件
	defer fileObj.Close()

	//读文件
	var tmp [128]byte
	for {
		n, err := fileObj.Read(tmp[:])
		if err == io.EOF {
			fmt.Println("读完了")
			return
		}
		if err != nil {
			fmt.Printf("read from file failed, err:%v\n", err)
			return
		}
		fmt.Printf("读了%d个字节\n", n)
		fmt.Println(string(tmp[:n]))
		if n < 128 {
			return
		}
	}
}

//利用bufio这个包读取文件
func readFromFilebyBufio() {
	fileObj, err := os.Open("./main.go")
	if err != nil {
		fmt.Printf("open file failed, err:%v\n", err)
		return
	}

	//关闭文件
	defer fileObj.Close()

	//创建一个用来从文件中都内容的对象
	reader := bufio.NewReader(fileObj)
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			return
		}
		if err != nil {
			fmt.Printf("read line failed,err:v%\n", err)
			return
		}
		fmt.Print(line)
	}
}

func readFromFilebyIoutil() {
	ret, err := ioutil.ReadFile("./main.go")
	if err != nil {
		fmt.Printf("read file failed, err:%v\n", err)
		return
	}
	fmt.Println(string(ret))
}
func main() {
	//readFromFile1()
	//readFromFilebyBufio()
	readFromFilebyIoutil()
}
