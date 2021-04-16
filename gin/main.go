package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/peter-wins/studygo/gin/chapter01"
	"github.com/peter-wins/studygo/gin/chapter02"
)



func main(){
	r:=gin.Default()

	r.LoadHTMLGlob("template/**/*")	// 二级目录的模板指定
	//r.LoadHTMLGlob("template/**/**/*")	// 三级目s录的模板指定
	//r.LoadHTMLFiles("template/index.html")

	r.GET("/", chapter01.Hello)
	r.GET("/user", chapter02.User)
	r.GET("arr", chapter02.ArrController)
	r.GET("/arr_struct",chapter02.ArrStruct)
	r.GET("/map", chapter02.MapController)
	r.GET("/map_struct", chapter02.MapStruct)
	r.GET("/slice", chapter02.SliceController)

	//在浏览器访问http://127.0.0.1:8080/param1/1242
	r.GET("/param1/:id",chapter02.Param1)

	//在浏览器访问http://127.0.0.1:8080/param2
	r.GET("/param2/*id",chapter02.Param2)

	//在浏览器访问http://127.0.0.1:8080/query?name=xxx
	r.GET("query",chapter02.GetQueryData)

	//在浏览器访问http://127.0.0.1:8080/query_arr/?ids=12,34,56
	r.GET("/query_arr",chapter02.GetQueryArrData)

	//在浏览器访问http://127.0.0.1:8080/query_map?user[name]=xxx&user[age]=19
	r.GET("/query_map",chapter02.GetQueryMapData)

	r.GET("/to_user_add",chapter02.ToUserAdd)
	r.POST("/do_user_add",chapter02.DoUserAdd)

	r.GET("/to_user_add2",chapter02.ToUserAdd2)
	r.POST("/do_user_add2",chapter02.DoUserAdd2)

	r.GET("/to_user_add3",chapter02.ToUserAdd3)
	r.POST("/do_user_add3",chapter02.DoUserAdd3)
	//文件上传
	r.GET("/to_upload1",chapter02.ToUpload1)
	r.POST("/do_upload1",chapter02.DoUpload1)
	//多文件上传
	r.GET("/to_upload2",chapter02.ToUpload2)
	r.POST("/do_upload2",chapter02.DoUpload2)

	//ajax文件上传
	r.GET("/to_upload3",chapter02.ToUpload3)
	r.POST("/do_upload3",chapter02.DoUpload3)

	err := r.Run(":8080")
	if err != nil {
		fmt.Println("start failed")
	}
}
