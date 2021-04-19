package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
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

	db.AutoMigrate(&models.User{})

	// 新增数据
	//db.Create(&models.User{Name: "bgd",Age: 28,Addr: "yxy",Pic: "/static/img/gg.png",Phone: 354533324})

	//查询数据
	//var user models.User
	//db.First(&user,3)	//"1"默认为id字段
	//db.First(&user, "name=?","bgd")	//指定字段查询
	//fmt.Println(user)

	// 更新数据
	//var user models.User
	//db.First(&user,1)
	//db.Model(&user).Update("age",22)
	//db.Model(&user).Update("name","xiaosb")
	//db.Model(&user).Update("addr","wes.ada.dubai")

	//删除数据
	var user models.User
	db.First(&user,3)
	db.Delete(&user)
}
