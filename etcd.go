package dsg

import "github.com/xgpc/dsg/pkg/etcd"

var _etcd *etcd.Handler

func OptionEtcd(conf etcd.Config) func() error {
	return func() error {

		_etcd = etcd.New(conf)
		if _etcd == nil {
			panic("start etcd error")
		}
		return nil
	}
}
