package main

import (
	"fmt"
	"net"
)

// tcp server端

func processConn(conn net.Conn) {
	// 3、与客户端通信
	defer conn.Close()
	var tmp [128]byte
	for {
		n, err := conn.Read(tmp[:])
		if err != nil {
			fmt.Println("read from conn failed, err:", err)
			return
		}
		fmt.Println(string(tmp[:n]))
	}
}
func main() {
	// 1、本地端口启动服务
	listener, err := net.Listen("tcp", "127.0.0.1:20000")
	if err != nil {
		fmt.Println("start tcp server on 127.0.0.1:20000 failed, err:%c", err)
		return
	}
	// 2、等待客户端建立连接
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("accept failed, err:", err)
			return
		}
		go processConn(conn)
	}
}
