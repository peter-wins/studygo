package router

import (
	"github.com/gin-gonic/gin"
	"github.com/peter-wins/studygo/gin/chapter01"
	"github.com/peter-wins/studygo/gin/chapter02"
	"github.com/peter-wins/studygo/gin/chapter03"
)

func Routers(r *gin.Engine){
	chap01 := r.Group("/chapter01")
	chap02 := r.Group("/chapter02")
	chap03 := r.Group("/chapter03")

	chapter01.Router(chap01)
	chapter02.Router(chap02)
	chapter03.Router(chap03)
}