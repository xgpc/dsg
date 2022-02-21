package router

import (
	"github.com/kataras/iris/v12"
	"github.com/xgpc/dsg/api/cryptExam"
)

func Json(party iris.Party) {
	r := party.Party("/json")
	//Sys
	r.Post("/send", cryptExam.SendJson)
	r.Get("/receive", cryptExam.BackJson)
}
