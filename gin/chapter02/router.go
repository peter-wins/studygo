package chapter02

import "github.com/gin-gonic/gin"

func Router(r *gin.RouterGroup){
	r.GET("/user", User)
	r.GET("arr", ArrController)
	r.GET("/arr_struct",ArrStruct)
	r.GET("/map", MapController)
	r.GET("/map_struct", MapStruct)
	r.GET("/slice", SliceController)

	//在浏览器访问http://127.0.0.1:8080/param1/1242
	r.GET("/param1/:id",Param1)

	//在浏览器访问http://127.0.0.1:8080/param2
	r.GET("/param2/*id",Param2)

	//在浏览器访问http://127.0.0.1:8080/query?name=xxx
	r.GET("query",GetQueryData)

	//在浏览器访问http://127.0.0.1:8080/query_arr/?ids=12,34,56
	r.GET("/query_arr",GetQueryArrData)

	//在浏览器访问http://127.0.0.1:8080/query_map?user[name]=xxx&user[age]=19
	r.GET("/query_map",GetQueryMapData)

	r.GET("/to_user_add",ToUserAdd)
	r.POST("/do_user_add",DoUserAdd)

	r.GET("/to_user_add2",ToUserAdd2)
	r.POST("/do_user_add2",DoUserAdd2)

	r.GET("/to_user_add3",ToUserAdd3)
	r.POST("/do_user_add3",DoUserAdd3)
	//文件上传
	r.GET("/to_upload1",ToUpload1)
	r.POST("/do_upload1",DoUpload1)
	//多文件上传
	r.GET("/to_upload2",ToUpload2)
	r.POST("/do_upload2",DoUpload2)
	//ajax文件上传
	r.GET("/to_upload3",ToUpload3)
	r.POST("/do_upload3",DoUpload3)
}