// Description: etcd demo
package main

import (
	"fmt"
	"github.com/xgpc/dsg/v2"
)

func EtcdDemo() {
	dsg.Load("config.yml")
	fmt.Println(dsg.Conf)
	dsg.Default(dsg.OptionEtcd(dsg.Conf.Etcd))
	dsg.Etcd.RegisterServiceDefault()
}
