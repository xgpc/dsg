package sys

import (
	"example/models"
	"github.com/kataras/iris/v12"
	"github.com/xgpc/dsg/frame"
)

// Dict 字典值
// @Summary      字典值
// @Description 字典值
// @Tag sys
// @Produce json
// @Success 200 {object} render.Response
// @Router /api/sys/dict [get]
func Dict(ctx iris.Context) {
	this := frame.NewBase(ctx)
	this.SuccessWithData(models.Dict)
}
