// Package dsg
// @Author:        asus
// @Description:   $
// @File:          New
// @Data:          2022/2/2118:09
//
package dsg

import (
	"github.com/kataras/iris/v12"
	"github.com/xgpc/dsg/frame"
	"github.com/xgpc/dsg/models"
	"github.com/xgpc/dsg/service/cryptService"
	"github.com/xgpc/dsg/service/grpcService/proto"
	"github.com/xgpc/dsg/service/schedule"
	"github.com/xgpc/dsg/service/sysService"
	"github.com/xgpc/dsg/service/validatorService"
	path2 "path"
)

type Service struct {
	App *iris.Application
}

func New(paths ...string) *Service {
	app := iris.New()

	path := path2.Join(paths...)
	if path == "" {
		path = "config.yaml"
	}

	// 加载配置
	frame.Load(app, path)

	// 加载mysql
	if frame.Config.Mysql.Host != "" {
		frame.LoadMysql()
	}

	// 加载Redis
	if frame.Config.Redis.Host != "" {
		frame.LoadRedis()
	}

	// 加载服务
	if frame.Config.SysConfig.StartSchedule {
		schedule.Start()
		schedule.StartSchedules()
	}

	// 参数验证器配置
	if frame.Config.SysConfig.ValidatorService {
		validatorService.GetTranslations()
	}

	//	系统时间版本
	sysService.InitSysVersion()

	//	默认user表
	if frame.Config.SysConfig.UserDefault {
		models.InitUser()
	}

	// rsa秘钥初始化
	if frame.Config.SysConfig.GenerateRSAKey {
		cryptService.GenerateRSAKey(1024)
		cryptService.SetRsaKey()
	}

	// grpc连接初始化
	if frame.Config.Microservices.RPCAddr != "" {
		proto.GRPCConnect()
	}

	return &Service{
		App: app,
	}
}

// Start 启动监听
func (app *Service) Start() {
	frame.Listening(app.App)
}
