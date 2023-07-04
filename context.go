package dsg

import (
	"github.com/kataras/iris/v12"
)

type Context struct {
	Ctx iris.Context
}

func New(ctx iris.Context) *Context {

	return &Context{Ctx: ctx}
}
