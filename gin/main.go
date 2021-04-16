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

	err := r.Run(":8080")
	if err != nil {
		fmt.Println("start failed")
	}
}
