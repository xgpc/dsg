// Package fileService
// @Author:        asus
// @Description:   $
// @File:          uploadController
// @Data:          2022/3/2911:17
//
package fileService

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"github.com/xgpc/dsg/exce"
	"github.com/xgpc/dsg/frame"
	"github.com/xgpc/dsg/util/guzzle"
	"net/http"
)

type UploadController struct {
	Frame *frame.Base
	Ctx   iris.Context
}

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

var uploadClient *guzzle.Client

func (this *UploadController) BeforeActivation(b mvc.BeforeActivation) {
	b.Handle(http.MethodPost, "/oss/isExist", "OssExist")
}

// PostConfig 获取上传文件配置信息
// @Summary      获取上传文件配置信息
// @Description  获取上传文件配置信息
// @Produce      json
// @Tags         上传文件
// @param        token  header    string  true  "登录用户token"
// @Success      200    {object}  Response
// @Router       /api/server/config [post]
func (this *UploadController) PostConfig() {
	resp, err := uploadClient.RequestJSON(map[string]interface{}{}).Post("server/config")
	if err != nil {
		exce.ThrowSys(exce.CodeRequestError, err.Error())
	}

	res := Response{}
	resp.JSON(&res)
	if res.Code != 0 {
		exce.ThrowSys(exce.CodeRequestError, res.Msg)
	}
	this.Frame.SuccessWithData(res.Data)
}

type PostOssUpload struct {
	ProjectName string `valid:"required" json:"project_name"`
	FileType    string `valid:"required" json:"file_type"`
	FileSize    int64  `valid:"required" json:"file_size"`
}

// PostOssUpload 上传
// @Summary      上传
// @Description  上传
// @Produce      json
// @Tags         上传文件
// @param        token  header    string         true  "登录用户token"
// @param        root   body      PostOssUpload  true  "上传参数"
// @Success      200    {object}  Response
// @Router       /api/server/oss/upload [post]
func (this *UploadController) PostOssUpload() {
	param := &PostOssUpload{}
	this.Frame.Init(param)

	resp, err := uploadClient.RequestJSON(map[string]interface{}{
		"ProjectName": param.ProjectName,
		"FileType":    param.FileType,
		"FileSize":    param.FileSize,
	}).Post("/server/oss/upload")
	if err != nil {
		exce.ThrowSys(exce.CodeRequestError, err.Error())
	}

	res := Response{}
	resp.JSON(&res)
	if res.Code != 0 {
		exce.ThrowSys(exce.CodeRequestError, res.Msg)
	}
	this.Frame.SuccessWithData(res.Data)
}

type PostOssConfirm struct {
	Key        string `valid:"required" json:"key"`
	ProductKey string `valid:"required" json:"product_key"`
}

// PostOssConfirm 上传确认
// @Summary      上传确认
// @Description  上传确认
// @Produce      json
// @Tags         上传文件
// @param        token  header    string          true  "登录用户token"
// @param        root   body      PostOssConfirm  true  "上传参数"
// @Success      200    {object}  Response
// @Router       /api/server/oss/confirm [post]
func (this *UploadController) PostOssConfirm() {
	param := &PostOssConfirm{}
	this.Frame.Init(param)

	resp, err := uploadClient.RequestJSON(map[string]interface{}{
		"Key":        param.Key,
		"ProductKey": param.ProductKey,
	}).Post("/server/oss/confirm")
	if err != nil {
		exce.ThrowSys(exce.CodeRequestError, err.Error())
	}

	res := Response{}
	resp.JSON(&res)
	if res.Code != 0 {
		exce.ThrowSys(exce.CodeRequestError, res.Msg)
	}
	this.Frame.SuccessWithData(res.Data)
}

type PostOssKey struct {
	Key string `valid:"required" json:"Key"`
}

// PostOssDelete 删除
// @Summary      删除
// @Description  删除
// @Produce      json
// @Tags         上传文件
// @param        token  header    string      true  "登录用户token"
// @param        root   body      PostOssKey  true  "上传参数"
// @Success      200    {object}  Response
// @Router       /api/server/oss/delete [post]
func (this *UploadController) PostOssDelete() {
	param := &PostOssKey{}
	this.Frame.Init(param)

	resp, err := uploadClient.RequestJSON(map[string]interface{}{
		"Key":    param.Key,
		"UserID": this.Frame.MyId(),
	}).Post("/server/oss/delete")
	if err != nil {
		exce.ThrowSys(exce.CodeRequestError, err.Error())
	}

	res := Response{}
	resp.JSON(&res)
	if res.Code != 0 {
		exce.ThrowSys(exce.CodeRequestError, res.Msg)
	}
	this.Frame.SuccessWithData(res.Data)
}

// OssExist 查看是否存在
// @Summary      查看是否存在
// @Description  查看是否存在
// @Produce      json
// @Tags         上传文件
// @param        token  header    string      true  "登录用户token"
// @param        root   body      PostOssKey  true  "上传参数"
// @Success      200    {object}  Response
// @Router       /api/server/oss/isExist [post]
func (this *UploadController) OssExist() {
	param := &PostOssKey{}
	this.Frame.Init(param)

	resp, err := uploadClient.RequestJSON(map[string]interface{}{
		"Key": param.Key,
	}).Post("/server/oss/isExist")
	if err != nil {
		exce.ThrowSys(exce.CodeRequestError, err.Error())
	}

	res := Response{}
	resp.JSON(&res)
	if res.Code != 0 {
		exce.ThrowSys(exce.CodeRequestError, res.Msg)
	}
	this.Frame.SuccessWithData(res.Data)
}
