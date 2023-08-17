package main

import (
	"fmt"
	"github.com/xgpc/dsg/v2"
)

func main() {
	// dsg 使用
	dsg.Load("config.yml")
	fmt.Printf("%+v", dsg.Conf)
}
