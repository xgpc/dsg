package dsg

import "github.com/xgpc/dsg/v2/pkg/etcd"

var Etcd *etcd.Handler

func GetServiceList(serverName string) []etcd.Service {
	services, err := Etcd.DiscoverServices(serverName)
	if err != nil {
		panic(serverName + "未能找到服务:" + err.Error())
	}

	return services
}

func OptionEtcd(conf etcd.Config) func() error {
	return func() error {

		Etcd = etcd.New(conf)
		if Etcd == nil {
			panic("start etcd error")
		}
		return nil
	}
}
