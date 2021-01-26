package main

import (
	"fmt"
	"os"
)

func f1() {
	var fileObj *os.File
	var err error
	fileObj, err := os.Open("./main.go")
	defer fileObj.Close() //当ERR有值的时候，fileObj就是nil，nil不能调用Close()
	if err != nil {
		fmt.Println("open file failed, err%v", err)
		return
	}
}
func f2() {
	var fileObj *os.File
	var err error
	fileObj, err := os.Open("./main.go")
	defer fileObj.Close()
	if err != nil {
		fmt.Println("open file failed, err%v", err)
		return
	}
	defer fileObj.Close()
}
func main() {
	f2()
}
