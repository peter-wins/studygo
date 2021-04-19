package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/peter-wins/studygo/ginLogin_project/controllers"
)

func InitRouter()*gin.Engine{
	router := gin.Default()
	router.LoadHTMLGlob("views/*")
	//注册
	router.GET("/register",controllers.RegisterGet)
	return router
}