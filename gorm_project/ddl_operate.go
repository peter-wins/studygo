package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/peter-wins/studygo/gorm_project/models"
)


func main(){
	//用户名:密码@tcp(ip:port)/数据库?charset=utf-8&parseTime=True&loc=Local
	db, err := gorm.Open("mysql",
		"root:73da0b2bf5e36958@tcp(192.192.191.244:3306)/gorm_project?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
	panic(err)
	}

	defer db.Close()
	////建表
	//db.CreateTable(&User{})	// 表users
	////指定表名
	//db.Table("user").CreateTable(&User{})
	////删除表
	//db.DropTable("user")
	//db.DropTable(&User{})
	////检查表是否存在
	//b := db.HasTable("user")
	//b1 := db.HasTable(&User{})	//判断users表是否存在
	//fmt.Println(b)
	//fmt.Println(b1)

	//自动迁移
	db.AutoMigrate(&models.User{},&models.UserInfo{})
}

