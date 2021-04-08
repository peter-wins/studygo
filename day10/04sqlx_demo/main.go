package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB

type user struct {
	ID   int
	Name string
	Age  int
}

func initDB() (err error) {
	dsn := "root:73da0b2bf5e36958@tcp(192.192.191.244:3306)/sql_test"
	db, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		return
	}
	return
}
func main() {
	err := initDB()
	if err != nil {
		fmt.Printf("Init DB failed, err:%v\n", err)
	}
	fmt.Println("连接数据库成功!")

	sqlStr1 := `select id, name, age from user where id=7;`
	var u user
	db.Get(&u, sqlStr1)
	fmt.Printf("u:%#v\n", u)

	// 查询多条数据
	var userList []user
	sqlStr2 := `select id, name, age from user`
	err = db.Select(&userList, sqlStr2)
	if err != nil {
		fmt.Printf("select failed, err:%v\n", err)
		return
	}
	fmt.Printf("userList: %#v\n", userList)

}
