package chapter03
import (
"fmt"
"github.com/gin-gonic/gin"
"net/http"
)

func BindQueryString(ctx *gin.Context){
	var p2 User
	err := ctx.ShouldBind(&p2)	//ShouldBind绑定
	if err != nil {
		ctx.String(http.StatusNotFound,"绑定失败")
	}else {
		ctx.String(http.StatusOK,"绑定成功")
		fmt.Println(p2)
	}
}
