package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/peter-wins/studygo/gin/router"
)

func main() {
	r := gin.Default()
	//创建日志文件,日志写入文件
	//file, _ := os.Create("gin.log")
	//gin.DefaultWriter = io.MultiWriter(file,os.Stdout)

	r.LoadHTMLGlob("template/**/*") // 二级目录的模板指定
	//r.LoadHTMLGlob("template/**/**/*")	// 三级目s录的模板指定
	//r.LoadHTMLFiles("template/index.html")
	r.Static("/static", "static")
	//r.StaticFS("/static",http.Dir("static"))

	//路由组示例
	//v1 := r.Group("/v1")
	//{
	//	v1.GET("/", chapter01.Hello)
	//}
	//
	router.Routers(r)

	err := r.Run(":8080")
	if err != nil {
		fmt.Println("start failed")
	}
}
