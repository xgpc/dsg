package router

import (
	"github.com/kataras/iris/v12"
	"github.com/xgpc/dsg/api/cryptExam"
)

func MsgCode(party iris.Party) {
	r := party.Party("/msg/code")
	//Sys
	r.Post("/", cryptExam.SendCode)
}
