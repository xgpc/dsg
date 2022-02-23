package main

import (
	"github.com/xgpc/dsg"
)

func main() {
	server := dsg.New(dsg.SysConfig{})

	//server.App

	server.Start()
}
