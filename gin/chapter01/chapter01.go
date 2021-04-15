package chapter01

import "github.com/gin-gonic/gin"

func Hello (ctx *gin.Context){
	name:= "李斯"
	ctx.HTML(200, "index/index.html", name)
}