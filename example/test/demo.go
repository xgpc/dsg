package main

import (
	"fmt"
	"github.com/xgpc/dsg/v2"
)

func main() {
	dsg.Load("config.yml")
	fmt.Println(dsg.Conf)
}
