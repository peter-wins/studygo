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

	var person UserInfo
	person.Id = 2
	person.Name = "十八岁的小姐姐"
	person.Age = 18
	person.Addr = "荷兰"

	ctx.HTML(http.StatusOK,"chapter02/user.html", person)
}
// ArrController 渲染数组
func ArrController(ctx *gin.Context){
	arr := [3]int{1,3,4}	// 定义一个数组
	ctx.HTML(http.StatusOK,"chapter02/arr.html", arr)
}
// ArrStruct 结构体数组渲染
func ArrStruct(ctx *gin.Context){
	arrStruct := [3]UserInfo{
		{Id: 1, Name: "娜娜", Age: 19, Addr: "新疆"},
		{Id: 2, Name: "莉莉", Age: 29, Addr: "上海"},
		{Id: 3, Name: "等等", Age: 30, Addr: "北京"},
	}
	ctx.HTML(http.StatusOK,"chapter02/arr_struct.html", arrStruct)
}

// MapController map 渲染
func MapController(ctx *gin.Context){
	m1 := map[string]string{"name": "玥儿", "age": "19"}
	m2 := map[string]int{"age": 20}
	m3 := map[string]interface{}{"m1":m1, "m2":m2}
	ctx.HTML(http.StatusOK, "chapter02/map.html", m3)
}

// MapStruct map+struct渲染
func MapStruct(ctx *gin.Context){
	mapStruct := map[string]UserInfo{
		"p1":{Id: 1, Name: "嬛嬛", Age: 29, Addr: "上古石窟"},
		"p2":{Id: 2, Name: "九九", Age: 18, Addr: "台北花园"},
		"p3":{Id: 3, Name: "龙盛", Age: 26, Addr: "青斑石门"},
		"p4":{Id: 4, Name: "吉吉国王", Age: 39, Addr: "原始森林"},
	}
	ctx.HTML(http.StatusOK,"chapter02/map_struct.html", mapStruct )
}
// slice渲染
func SliceController (ctx *gin.Context){
	s1 := []int{1,2,3,4,5,6,7,8}
	ctx.HTML(http.StatusOK,"chapter02/slice.html",s1)
}