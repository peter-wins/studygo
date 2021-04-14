package main

import (
	"context"
	"fmt"
	"go.etcd.io/etcd/clientv3"
	"time"
)

func main() {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"192.192.191.244:2379"},
		DialTimeout: 2 * time.Second,
	})
	if err != nil {
		fmt.Printf("connect to ectd failed, err:%v\n", err)
		return
	}
	fmt.Println("connect to etcd success~~")

	defer cli.Close()

	// put
	ctx, cancel:=context.WithTimeout(context.Background(), time.Second)
	_,err = cli.Put(ctx, "xiaohua", "good girl")
	cancel()
	if err != nil{
		fmt.Printf("put to etcd failed, err:%v\n", err)
		return
	}

	// get
	ctx, cancel =context.WithTimeout(context.Background(), time.Second)
	resp, err := cli.Get(ctx, "xiaohua")
	cancel()
	if err != nil {
		fmt.Printf("get from etcd failed, err:%v\n",err)
		return
	}
	for _, ev := range resp.Kvs{
		fmt.Printf("%s:%s\n", ev.Key,ev.Value)
	}
}