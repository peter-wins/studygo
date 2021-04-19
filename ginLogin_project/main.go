package main

import (
	"fmt"
	"github.com/peter-wins/studygo/ginLogin_project/routers"
)

func main(){
	router := routers.InitRouter()
	//静态资源
	router.Static("/static","static")
	err := router.Run(":8081")
	if err != nil {
		fmt.Printf("service start failed, err:%v\n",  err)
		return
	}
}
