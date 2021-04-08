package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

type user struct {
	id   int
	name string
	age  int
}

func initDB() (err error) {
	dsn := "root:73da0b2bf5e36958@tcp(192.192.191.244:3306)/sql_test"
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		return
	}
	err = db.Ping()
	if err != nil {
		return
	}
	return
}

func transactionDemo() {
	// 开启事务
	tx, err := db.Begin()
	if err != nil {
		fmt.Printf("begin failed, err%v\n", err)
		return
	}
	// 执行多个SQL语句
	sqlStr1 := `update user set age=age-2 where id = 1;`
	sqlStr2 := `update user set age=age+2 where id = 2;`

	_, err = tx.Exec(sqlStr1)
	if err != nil {
		// 要回滚
		tx.Rollback()
		fmt.Println("执行SQL1失败了,已回滚~")
		return
	}
	_, err = tx.Exec(sqlStr2)
	if err != nil {
		// 要回滚
		tx.Rollback()
		fmt.Println("执行SQL2失败了,已回滚~")
		return
	}
	// 提交事务
	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		fmt.Println("提交事务失败,已回滚~")
		return
	}
	fmt.Println("Exec trans success ~~")
}

func main() {
	err := initDB()
	if err != nil {
		fmt.Printf("Init DB failed, err:%v\n", err)
	}
	fmt.Println("连接数据库成功!")
	transactionDemo()
}
