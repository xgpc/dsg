package main

import (
	"fmt"
	"github.com/xgpc/dsg/v2/pkg/etcd"

	clientv3 "go.etcd.io/etcd/client/v3"
	"time"
)

func main() {
	// 链接etcd

	conf := etcd.Config{
		Name:                 "api-gateway",
		Address:              "127.0.0.1",
		Port:                 8082,
		Endpoints:            []string{"http://127.0.0.1:2379"},
		AutoSyncInterval:     0,
		DialTimeout:          10,
		DefLeaseSecond:       10,
		DialKeepAliveTime:    0,
		DialKeepAliveTimeout: 0,
	}

	client := etcd.New(conf)

	// 查询key
	get, err := client.Get("1", clientv3.WithPrefix())
	if err != nil {
		panic(err)
	}

	for _, v := range get.Kvs {
		fmt.Printf("key:%s, value:%s\n", v.Key, v.Value)
	}

	// 注册服务
	err = client.RegisterServiceDefault()
	if err != nil {
		panic(err)
	}

	// 发现服务
	for {
		services, err := client.DiscoverServices("/apps/")
		if err != nil {
			panic(err)
		}

		fmt.Println(services)
		time.Sleep(time.Second * 10)
	}

	select {}

}
