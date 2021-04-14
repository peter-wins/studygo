package main

import (
	"context"
	"fmt"
	"github.com/olivere/elastic/v7"
)
// ES insert data demo

type Student struct {
	Name string	`json:"name"`
	Age int	`json:"age"`
	Married bool `json:"married"`
}
func main() {
	//1. 初始化连接，得到一个client
	client,err:=elastic.NewClient(elastic.SetURL("http://192.192.191.244:9200"))
	if err != nil {
		panic(err)
	}
	println("connect to elasticsearch success")

	p1 := Student{Name: "Rion", Age: 22, Married: false}
	// 链式操作
	put1,err := client.Index().Index("student").BodyJson(p1).Do(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Printf("Indexed user %s to index %s, type %s\n", put1.Id, put1.Index, put1.Type)
}
