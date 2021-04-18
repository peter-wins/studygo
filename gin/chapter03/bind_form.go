package chapter03

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)
type User struct {
	Name string	`form:"name"`
	Age int `form:"age"`
	Addr string `form:"addr"`
}
func ToBindForm(ctx *gin.Context){
	ctx.HTML(http.StatusOK,"chapter03/bind_form.html",nil)
}
func DoBindForm(ctx *gin.Context){
	var p1 User
	err := ctx.ShouldBind(&p1)	//ShouldBind绑定
	if err != nil {
		ctx.String(http.StatusNotFound,"绑定失败")
	}else {
		ctx.String(http.StatusOK,"绑定成功")
		fmt.Println(p1)
	}
}