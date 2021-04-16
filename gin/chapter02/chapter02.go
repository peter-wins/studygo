package chapter02

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserInfo struct {
	Id int	`form:"id"`
	Name string	`form:"name"`
	Age int	`form:"age"`
	Addr string	`form:"addr"`
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
// SliceController slice渲染
func SliceController (ctx *gin.Context){
	s1 := []int{1,2,3,4,5,6,7,8}
	ctx.HTML(http.StatusOK,"chapter02/slice.html",s1)
}
// Param1 路径中直接加上参数值
func Param1(ctx *gin.Context){
	id := ctx.Param("id")
	ctx.String(http.StatusOK,"Hello: %s",id)
}
func Param2(ctx *gin.Context){
	id := ctx.Param("id")
	ctx.String(http.StatusOK,"Hello: %s",id)
}
func GetQueryData(ctx *gin.Context){
	id := ctx.Query("id")
	name := ctx.DefaultQuery("name","阿三")
	fmt.Println(name)
	ctx.String(http.StatusOK, "Hello: %s", id)
}
func GetQueryArrData(ctx *gin.Context){
	ids := ctx.QueryArray("ids")
	ctx.String(http.StatusOK, "Hello: %s", ids)
}
func GetQueryMapData(ctx *gin.Context){
	user := ctx.QueryMap("user")
	ctx.String(http.StatusOK, "Hello: %s", user)
}
// ToUserAdd POST请求(用户添加页面)
func ToUserAdd(ctx *gin.Context){
	ctx.HTML(http.StatusOK,"chapter02/user_add.html",nil)
}
// DoUserAdd 添加用户
func DoUserAdd(ctx *gin.Context){
	// PostForm用法
	userName := ctx.PostForm("username")
	passWord := ctx.PostForm("password")
	// DefaultPostForm用法
	age := ctx.DefaultPostForm("age","19")
	// PostFormArray用法
	loves := ctx.PostFormArray("love")
	// PostFormMap用法
	users := ctx.PostFormMap("user")

	fmt.Println(userName)
	fmt.Println(passWord)
	fmt.Println(loves)
	fmt.Println(users)
	fmt.Println(age)

	ctx.String(http.StatusOK, "添加成功~")
}
// ToUserAdd2 Ajax获取post请求数据
func ToUserAdd2(ctx *gin.Context){
	ctx.HTML(http.StatusOK,"chapter02/user_add2.html",nil)
}

func DoUserAdd2(ctx *gin.Context) {
	userName := ctx.PostForm("username")
	passWord := ctx.PostForm("passWord")
	age := ctx.PostForm("age")

	fmt.Println(userName)
	fmt.Println(passWord)
	fmt.Println(age)

	//m := map[string]interface{}{
	//	"code":200,
	//	"msg":"添加成功~",
	//}
	//
	//ctx.JSON(http.StatusOK,m)
	ctx.JSON(http.StatusOK,gin.H{"code":200,"msg":"添加成功~"})
}
func ToUserAdd3(ctx *gin.Context){
	ctx.HTML(http.StatusOK,"chapter02/user_add3.html",nil)
}
func DoUserAdd3(ctx *gin.Context) {

	var p2 UserInfo
	err := ctx.ShouldBind(&p2)
	fmt.Println(err)
	fmt.Println(p2)

	ctx.String(http.StatusOK,"绑定成功")
}