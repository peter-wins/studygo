package main

import (
	"fmt"
	"net/url"
)

//写一个整点报时程序，每当整点的时候发送一条Telegram群消息
//1、定时器
//13:00:00 12:00:00
//2、执行发送图片命令

// Manage the HTTP GET request parameters
type GetRequest struct {
	urls url.Values
}

// Initializer
func (p *GetRequest) Init() *GetRequest {
	p.urls = url.Values{}
	return p
}

// Initialized from another instance
func (p *GetRequest) InitFrom(reqParams *GetRequest) *GetRequest {
	if reqParams != nil {
		p.urls = reqParams.urls
	} else {
		p.urls = url.Values{}
	}
	return p
}

// Add URL escape property and value pair
func (p *GetRequest) AddParam(property string, value string) *GetRequest {
	if property != "" && value != "" {
		p.urls.Add(property, value)
	}
	return p
}

// Concat the property and value pair
func (p *GetRequest) BuildParams() string {
	return p.urls.Encode()
}

func main() {
	init := new(GetRequest).Init()
	params := init.AddParam("market", "sh").AddParam("Inst", "6000987").BuildParams()
	//params := init.AddParam("market", "sh")
	fmt.Println(params)
}
