package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql" // init()
)

var db *sql.DB //是一个数据库连接池对象

func initDB() (err error) {
	// 数据库信息
	dsn := "root:73da0b2bf5e36958@tcp(192.192.191.244:3306)/sql_test"
	// 连接数据库
	db, err = sql.Open("mysql", dsn) // 这里要用 "=" 不能用 ":="
	if err != nil {
		return
	}
	err = db.Ping()
	if err != nil {
		return
	}
	// 设置数据库连接池的最大连接数
	db.SetMaxOpenConns(10)
	// 设置数据库连接池最大空闲连接数
	db.SetMaxIdleConns(5)
	return
}

type user struct {
	id   int
	name string
	age  int
}

// 查询单行记录
func queryOne(id int) {
	var u1 user
	// 1.写查询单挑记录的sql语句
	sqlStr := `select id,name,age from user where id=?;` // ？占位符
	// 2.执行
	rowObj := db.QueryRow(sqlStr, id) // 从连接池里拿一个出来去数据库查询单条记录
	// 3.拿到结果
	rowObj.Scan(&u1.id, &u1.name, &u1.age)
	//rowObj := db.QueryRow(sqlStr, 1).Scan(&u1.id, &u1.name, &u1.age) // 简洁写法
	// 4.打印结果
	fmt.Printf("u1:%#v\n", u1)
}

// 插入
func insert() {

}

func main() {
	err := initDB()
	if err != nil {
		fmt.Printf("init DB failed, err:%v\n", err)
	}
	fmt.Println("连接数据库成功！")
	queryOne(2)

}
