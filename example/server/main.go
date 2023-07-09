package main

import "github.com/xgpc/dsg/v2/pkg/etcd"

func main() {
	conf := etcd.Config{
		Name:                 "/apps/{server-name}",
		Address:              "127.0.0.1",
		Port:                 8081,
		Endpoints:            []string{"http://127.0.0.1:2379"},
		AutoSyncInterval:     0,
		DialTimeout:          10,
		DefLeaseSecond:       10,
		DialKeepAliveTime:    0,
		DialKeepAliveTimeout: 0,
	}
	client := etcd.New(conf)

	err := client.RegisterServiceDefault()
	if err != nil {
		panic(err)
	}

	select {}
}
