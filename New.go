// Package dsg
// @Author:        asus
// @Description:   $
// @File:          New
// @Data:          2022/2/2118:09
//
package dsg

import (
	"github.com/xgpc/dsg/env"
	path2 "path"

	"github.com/kataras/iris/v12"
	"github.com/xgpc/dsg/service/cryptService"
	"github.com/xgpc/dsg/service/grpcService/proto"
	"github.com/xgpc/dsg/service/schedule"
	"github.com/xgpc/dsg/service/sysService"
	"github.com/xgpc/dsg/service/validatorService"
)

type Service struct {
	App *iris.Application
}

func New(paths ...string) *Service {
	app := iris.New()

	path := path2.Join(paths...)
	if path == "" {
		path = "config.yml"
	}

	// 加载配置
	Load(app, path)

	// 加载mysql
	if env.Config.Mysql.Host != "" {
		LoadMysql()
	}

	// 加载Redis
	if env.Config.Redis.Host != "" {
		LoadRedisDefault()
	}

	// 加载服务
	if env.Config.SysConfig.StartSchedule {
		schedule.Start()
		schedule.StartSchedules()
	}

	// 参数验证器配置
	if env.Config.SysConfig.ValidatorService {
		validatorService.GetTranslations()
	}

	//	系统时间版本
	sysService.InitSysVersion()

	// rsa秘钥初始化
	if env.Config.SysConfig.GenerateRSAKey {
		cryptService.GenerateRSAKey(1024)
		cryptService.SetRsaKey()
	}

	// grpc连接初始化
	if env.Config.Microservices.RPCAddr != "" {
		proto.GRPCConnect()
	}

	return &Service{
		App: app,
	}
}

// Start 启动监听
func (app *Service) Start() {
	Listening(app.App)
}
