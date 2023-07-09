package dsg

import "github.com/xgpc/dsg/v2/pkg/etcd"

var _etcd *etcd.Handler

func GetServiceList(serverName string) []etcd.Service {
	services, err := _etcd.DiscoverServices(serverName)
	if err != nil {
		panic(serverName + "未能找到服务:" + err.Error())
	}

	return services
}

func OptionEtcd(conf etcd.Config) func() error {
	return func() error {

		_etcd = etcd.New(conf)
		if _etcd == nil {
			panic("start etcd error")
		}
		return nil
	}
}
