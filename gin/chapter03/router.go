package chapter03

import (
	"github.com/gin-gonic/gin"
)

func Router(r *gin.RouterGroup){
	//数据绑定
	r.GET("/to_bind_form",ToBindForm)
	r.POST("/do_bind_form",DoBindForm)
	r.GET("/bind_query_string",BindQueryString)
}