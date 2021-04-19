package database

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var db *sql.DB

func InitMysql(){
	fmt.Println("Init Mysql...")
	if db == nil {
		db, _ := sql.Open("mysql","root:73da0b2bf5e36958@tcp(192.192.191.244:3306/blogweb_gin?charset=utf8)")
	}
	CreateTableWithUser()
}
//ModifyDB 操作数据库
func ModifyDB(sql string, args ...interface{})(int64, error){
	result, err := db.Exec(sql, args...)
	if err != nil {
		log.Println(err)
		return 0, err
	}
	count, err := result.RowsAffected()
	if err != nil {
		log.Println(err)
		return 0, err
	}
	return count, nil
}
//CreateTableWithUser 创建用户表
func CreateTableWithUser(){
	sql := `CREATE TABLE IF NOT EXISTS users(
		id INT(4) PRIMARY KEY AUTO_INCREMENT NOT NULL,
		username VARCHAR(64),
		password VARCHAR(64),
		status INT(4),
		createtime INT(10)
		);`
	result, err := ModifyDB(sql)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(result)

}
//QueryRowDB 查询
func QueryRowDB(sql string) *sql.Row {
	return db.QueryRow(sql)
}