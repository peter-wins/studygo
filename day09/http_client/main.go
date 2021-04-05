package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	resp, err := http.Get("http://127.0.0.1:9090/hello/")
	if err != nil {
		fmt.Printf("get url failed, err:%v\n", err)
		return
	}
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("read resp.body failed, err:%v\n", err)
		return
	}
	fmt.Println(string(b))
}
