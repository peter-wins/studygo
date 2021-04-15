package chapter02

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserInfo struct {
	Id int
	Name string
	Age int
	Addr string
}
// User 结构体渲染
func User (ctx *gin.Context){

	// 初始化结构体的两种方式
	//user_info := UserInfo{Id: 1, Name: "佳佳", Age: 19, Addr: "wewstburry"}

	var user_info UserInfo
	user_info.Id = 2
	user_info.Name = "十八岁的小姐姐"
	user_info.Age = 18
	user_info.Addr = "荷兰"

	ctx.HTML(http.StatusOK,"chapter02/user.html", user_info)
}
// ArrController 渲染数组
func ArrController(ctx *gin.Context){
	arr := [3]int{1,3,4}	// 定义一个数组
	ctx.HTML(http.StatusOK,"chapter02/arr.html", arr)
}
// 结构体 数组渲染
func ArrStruct(ctx *gin.Context){
	arr_struct := [3]UserInfo{
		{Id: 1, Name: "娜娜", Age: 19, Addr: "新疆"},
		{Id: 2, Name: "莉莉", Age: 29, Addr: "上海"},
		{Id: 3, Name: "等等", Age: 30, Addr: "北京"},
	}
	ctx.HTML(http.StatusOK,"chapter02/arr_struct.html", arr_struct)
}