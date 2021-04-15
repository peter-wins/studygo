package chapter02

import "github.com/gin-gonic/gin"

func User (ctx *gin.Context){
	ctx.HTML(200,"user/user.html",nil)
}