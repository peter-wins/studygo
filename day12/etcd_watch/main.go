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
	// watch
	// 派一个哨兵 一直监视着 xiaohua 这个key的变化（新增 修改 删除）
	ch := cli.Watch(context.Background(),"xiaohua")
	// 尝试从通道取值（监视的信息）
	for wresp := range ch {
		for _, ev := range wresp.Events{
			fmt.Printf("Type:%v key:%v value:%v\n", ev.Type, string(ev.Kv.Key), string(ev.Kv.Value))
		}
	}

}
