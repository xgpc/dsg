package apps

import (
	"company/admin/models"
	"github.com/xgpc/dsg/exce"
)

func reloadRouter() {
	var list []models.App
	err := db().Model(&models.App{}).Find(&list).Error
	if err != nil {
		exce.ThrowSys(exce.CodeRequestError, err.Error())
	}

	md := map[string]string{}
	for _, v := range list {
		if v.RouterApp != "" {
			md[v.RouterApp] = v.RouterPath
		}
	}

	routers = md
}
