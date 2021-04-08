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
	sqlStr := `select id, name, age from user where id=?;` // ？占位符
	// 2.执行
	rowObj := db.QueryRow(sqlStr, id) // 从连接池里拿一个出来去数据库查询单条记录
	// 3.拿到结果
	rowObj.Scan(&u1.id, &u1.name, &u1.age)
	//rowObj := db.QueryRow(sqlStr, 1).Scan(&u1.id, &u1.name, &u1.age) // 简洁写法
	// 4.打印结果
	fmt.Printf("u1:%#v\n", u1)
}

// 多行查询
func querryMore(n int) {
	// SQL语句
	sqlStr := `select id, name, age from user where id > ?;`
	// 执行
	rows, err := db.Query(sqlStr, n)
	if err != nil {
		fmt.Printf("exec %s query failed,err:%v\n", sqlStr, err)
		return
	}
	// 关闭rows
	defer rows.Close()
	// 循环取值
	for rows.Next() {
		var u1 user
		err := rows.Scan(&u1.id, &u1.name, &u1.age)
		if err != nil {
			fmt.Printf("scan failed, err%v\n", err)
			return
		}
		// 4.打印结果
		fmt.Printf("u1:%#v\n", u1)
	}
}

// 插入数据
func insert() {
	// 写SQL语句
	sqlStr := `insert into user(name,age) values("莉莉",30);`
	// 执行 exec
	ret, err := db.Exec(sqlStr)
	if err != nil {
		fmt.Printf("insert failed, err:%v\n", err)
		return
	}
	// 如果是插入数据的操作，能够拿到插入数据的id
	id, err := ret.LastInsertId()
	if err != nil {
		fmt.Printf("get id failed, err:%v\n", err)
		return
	}
	fmt.Printf("id:%d", id)
}

// 更新数据
func updateRow(newAge int, id int) {
	sqlStr := `update user set age=? where id=?;`
	ret, err := db.Exec(sqlStr, newAge, id)
	if err != nil {
		fmt.Printf("update data failed, err:%v\n", err)
		return
	}
	n, err := ret.RowsAffected()
	if err != nil {
		fmt.Printf("get id failed, err:%v\n", err)
		return
	}
	fmt.Printf("更新了%d行数据\n", n)
}

// 删除数据
func deleteData(id int) {
	sqlStr := `delete from user where id = ?;`
	ret, err := db.Exec(sqlStr, id)
	if err != nil {
		fmt.Printf("delete data falied, err%v\n", err)
		return
	}
	n, err := ret.RowsAffected()
	if err != nil {
		fmt.Printf("get id failed, err:%v\n", err)
		return
	}
	fmt.Printf("成功删除%d行数据\n", n)
}

// 预处理方式插入多条数据

func prepareInsert() {
	sqlStr := `insert into user(name,age) values(?,?);`
	stmt, err := db.Prepare(sqlStr)
	if err != nil {
		fmt.Printf("prepare failed, err:%v\n", err)
		return
	}
	defer stmt.Close()
	// 后续只需要拿到ｓｔｍｔ去执行一些操作
	var m = map[string]int{
		"小刘":    26,
		"小章":    36,
		"大郎":    20,
		"张三":    34,
		"猴子":    32,
		"八戒":    28,
		"三藏":    39,
		"小猪佩奇":  18,
		"龙五":    34,
		"诸葛":    25,
		"牛姐":    38,
		"集美":    26,
		"波多野结衣": 29,
	}
	for k, v := range m {
		stmt.Exec(k, v)
	}
}

func main() {
	err := initDB()
	if err != nil {
		fmt.Printf("Init DB failed, err:%v\n", err)
	}
	fmt.Println("连接数据库成功！")
	//queryOne(1)
	//querryMore(2)
	//insert()
	//updateRow(43, 4)
	// deleteData(6)
	prepareInsert()

}
