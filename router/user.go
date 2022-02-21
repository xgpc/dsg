package router

import (
	"github.com/kataras/iris/v12"
	"github.com/xgpc/dsg/api/cryptExam"
	"github.com/xgpc/dsg/api/user"
)

func User(party iris.Party) {
	r := party.Party("/user")
	//Sys
	r.Get("/{UserID}", cryptExam.BackMobile)
	r.Post("/login", user.Login)
}
