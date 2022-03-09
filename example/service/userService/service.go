package userService

import (
	"github.com/xgpc/dsg/example/models"
	"github.com/xgpc/dsg/frame"
)

type Method interface {
	Get(userID uint32) (info models.User, err error)
}

type service struct {
	frame.Base
}

func Service() *service {
	return &service{}
}

// Get 根据userID查询用户
func (u service) Get(userID uint32) (info models.User, err error) {
	//
	return info, nil
}
