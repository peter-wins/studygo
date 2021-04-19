package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func RegisterGet(c *gin.Context){
	//返回html
	c.HTML(http.StatusOK,"register.html",gin.H{"Title":"注册页"})
}