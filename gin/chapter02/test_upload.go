package chapter02

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)
// form表单单文件上传
func ToUpload1(ctx *gin.Context)  {
	ctx.HTML(http.StatusOK,"chapter02/test_upload1.html",nil)
}
func DoUpload1(ctx *gin.Context){
	//获取文件
	file, _ := ctx.FormFile("file")
	fmt.Println(file.Filename)
	//保存文件
	unitInt := time.Now().Unix()	// 获取当前时间
	timeUnitStr := strconv.FormatInt(unitInt, 10)	//转换当前时间为string类型

	filePath := "upload/" + timeUnitStr + file.Filename	//拼接文件路径
	err:= ctx.SaveUploadedFile(file, filePath)

	if err != nil {
		fmt.Printf("save file failed, err%v\n", err)
		return
	}
	ctx.String(http.StatusOK, "上传成功")
}
// form表单单文件上传
func ToUpload2(ctx *gin.Context)  {
	ctx.HTML(http.StatusOK,"chapter02/test_upload2.html",nil)
}
func DoUpload2(ctx *gin.Context){
	//获取文件
	form, _ := ctx.MultipartForm()// 多文件获取
	files := form.File["file"]
	for _, file := range files{
		fmt.Println(file.Filename)
		//保存文件
		unitInt := time.Now().Unix()	// 获取当前时间
		timeUnitStr := strconv.FormatInt(unitInt, 10)	//转换当前时间为string类型

		filePath := "upload/" + timeUnitStr + file.Filename	//拼接文件路径
		err:= ctx.SaveUploadedFile(file, filePath)

		if err != nil {
			fmt.Printf("save file failed, err%v\n", err)
			return
		}
	}
	ctx.String(http.StatusOK, "上传成功")
}
//ajax单文件上传
func ToUpload3(ctx *gin.Context)  {
	ctx.HTML(http.StatusOK,"chapter02/test_upload3.html",nil)
}
func DoUpload3(ctx *gin.Context){
	name := ctx.PostForm("name")
	fmt.Println(name)
	//获取文件
	file, _ := ctx.FormFile("file")
	fmt.Println(file.Filename)
	//保存文件
	unitInt := time.Now().Unix()	// 获取当前时间
	timeUnitStr := strconv.FormatInt(unitInt, 10)	//转换当前时间为string类型

	filePath := "upload/" + timeUnitStr + file.Filename	//拼接文件路径
	err:= ctx.SaveUploadedFile(file, filePath)

	if err != nil {
		fmt.Printf("save file failed, err%v\n", err)
		return
	}
	ctx.JSON(http.StatusOK,gin.H{"code":200,"msg":"添加成功~"})

}