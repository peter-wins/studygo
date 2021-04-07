package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql" // init()
)

func main() {
	// 数据库信息
	dsn := "root:73da0b2bf5e36958@tcp(192.192.191.244:3306)/day10"
	// 连接数据库
	db, err := sql.Open("mysql", dsn) // 不会校验用户名和密码是否正确
	if err != nil {
		fmt.Printf("dsn:%s invalid, err:%v\n", dsn, err)
		return
	}
	err = db.Ping()
	if err != nil {
		fmt.Printf("open %s failed, err:%v\n", dsn, err)
	}
	fmt.Println("连接数据库成功！")
}
