package controller

import (
	"data-manager/exterror"
	middlewares "data-manager/middleware"
	"data-manager/model"
	webCtx "data-manager/web/context"
	"log"
	"net/http"
	"os/exec"
	"path"
	"strings"
)

type PlatonecliRequest struct {
	Passphrase string `json:"passphrase" bson:"passphrase"`
}

var PlatonecliController = &PlatonecliRequest{}

func (p *PlatonecliRequest) GetAccount(ctx *webCtx.Context) {

	if err := ctx.ShouldBind(&p); nil != err {
		middlewares.ResponseError(ctx, http.StatusBadRequest, err, "参数绑定失败！")
		return
	}
	if p.Passphrase == "" {
		err := exterror.ErrParameterInvalid
		middlewares.ResponseError(ctx, http.StatusBadRequest, err, "Passphrase不能为空！")
		return

	}
	res, err := model.DefaultPlatonecliModel.GetAccount(ctx, p.Passphrase, p.Passphrase)
	if nil != err {
		middlewares.ResponseError(ctx, http.StatusBadRequest, err, "生成账户失败！")
		return
	}

	middlewares.ResponseSuccess(ctx, "添加用户成功", res)
}

func (p *PlatonecliRequest) GetSystemConfig(ctx *webCtx.Context) {
	res, err := model.DefaultPlatonecliModel.GetSystemConfig(ctx)
	if nil != err {
		middlewares.ResponseError(ctx, http.StatusBadRequest, err, "获取参数失败！")
		return
	}

	middlewares.ResponseSuccess(ctx, "获取参数成功", res)
}

func (p *PlatonecliRequest) SetSystemConfig(ctx *webCtx.Context) {

	var systemConfig model.SystemConfigBody
	if err := ctx.ShouldBind(&systemConfig); nil != err {
		middlewares.ResponseError(ctx, http.StatusBadRequest, err, "参数绑定失败！")
		return
	}

	res, err := model.DefaultPlatonecliModel.SetSystemConfig(ctx, systemConfig)
	if nil != err {
		middlewares.ResponseError(ctx, http.StatusBadRequest, err, "操作失败！")
		return
	}

	middlewares.ResponseSuccess(ctx, "操作成功", res)
}

func (p *PlatonecliRequest) SetCnsName(ctx *webCtx.Context) {

	var cnsRequestInfo model.CnsRequest
	if err := ctx.ShouldBind(&cnsRequestInfo); nil != err {
		middlewares.ResponseError(ctx, http.StatusBadRequest, err, "参数绑定失败！")
		return
	}
	if cnsRequestInfo.Name == "" || cnsRequestInfo.Version == "" || cnsRequestInfo.Address == "" {
		err := exterror.ErrParameterInvalid
		middlewares.ResponseError(ctx, http.StatusBadRequest, err, "必填字段不能为空！")
		return

	}
	res, err := model.DefaultPlatonecliModel.SetCnsName(ctx, cnsRequestInfo)
	if nil != err {
		middlewares.ResponseError(ctx, http.StatusBadRequest, err, "操作失败！")
		return
	}

	middlewares.ResponseSuccess(ctx, "操作成功", res)
}

func (p *PlatonecliRequest) CnsRedirect(ctx *webCtx.Context) {

	var cnsRequestInfo model.CnsRequest
	if err := ctx.ShouldBind(&cnsRequestInfo); nil != err {
		middlewares.ResponseError(ctx, http.StatusBadRequest, err, "参数绑定失败！")
		return
	}
	if cnsRequestInfo.Name == "" || cnsRequestInfo.Version == "" {
		err := exterror.ErrParameterInvalid
		middlewares.ResponseError(ctx, http.StatusBadRequest, err, "必填字段不能为空！")
		return

	}
	res, err := model.DefaultPlatonecliModel.CnsRedirect(ctx, cnsRequestInfo)
	if nil != err {
		middlewares.ResponseError(ctx, http.StatusBadRequest, err, "操作失败！")
		return
	}

	middlewares.ResponseSuccess(ctx, "操作成功", res)
}

func (p *PlatonecliRequest) OpenFireWall(ctx *webCtx.Context) {

	var FireWallRequestInfo model.FireWallRequest
	if err := ctx.ShouldBind(&FireWallRequestInfo); nil != err {
		middlewares.ResponseError(ctx, http.StatusBadRequest, err, "参数绑定失败！")
		return
	}
	if FireWallRequestInfo.Address == "" {
		err := exterror.ErrParameterInvalid
		middlewares.ResponseError(ctx, http.StatusBadRequest, err, "必填字段不能为空！")
		return

	}
	res, err := model.DefaultPlatonecliModel.OpenFireWall(ctx, FireWallRequestInfo)
	if nil != err {
		middlewares.ResponseError(ctx, http.StatusBadRequest, err, "操作失败！")
		return
	}

	middlewares.ResponseSuccess(ctx, "操作成功", res)
}

func (p *PlatonecliRequest) CloseFireWall(ctx *webCtx.Context) {

	var FireWallRequestInfo model.FireWallRequest

	if err := ctx.ShouldBind(&FireWallRequestInfo); nil != err {
		middlewares.ResponseError(ctx, http.StatusBadRequest, err, "参数绑定失败！")
		return
	}

	if FireWallRequestInfo.Address == "" {
		err := exterror.ErrParameterInvalid
		middlewares.ResponseError(ctx, http.StatusBadRequest, err, "必填字段不能为空！")
		return

	}
	res, err := model.DefaultPlatonecliModel.CloseFireWall(ctx, FireWallRequestInfo)
	if nil != err {
		middlewares.ResponseError(ctx, http.StatusBadRequest, err, "操作失败！")
		return
	}

	middlewares.ResponseSuccess(ctx, "操作成功", res)
}

func (p *PlatonecliRequest) SetVal(ctx *webCtx.Context) {

	var Data model.DataVal

	if err := ctx.ShouldBind(&Data); nil != err {
		middlewares.ResponseError(ctx, http.StatusBadRequest, err, "参数绑定失败！")
		return
	}

	params := []string{"invoke", "--addr", "0x1a34170b185c9fa29287adb8ffe363d69b97e27b", "--abi", "/home/colin/Workspace/src/sc/build/sc/sc.abi.json", "--config", "/home/colin/Workspace/src/PlatONE-Go/release/linux/conf/ctool.json", "--func", "set", "--param"}
	params = append(params, Data.Data)
	res, _ := runScript(params)
	middlewares.ResponseSuccess(ctx, "操作成功", res)
}

func (p *PlatonecliRequest) GetVal(ctx *webCtx.Context) {

	res, err := runScript([]string{"invoke", "--addr", "0x1a34170b185c9fa29287adb8ffe363d69b97e27b", "--abi", "/home/colin/Workspace/src/sc/build/sc/sc.abi.json", "--config", "/home/colin/Workspace/src/PlatONE-Go/release/linux/conf/ctool.json", "--func", "get"})
	if err != nil {
		middlewares.ResponseError(ctx, http.StatusBadRequest, err, "报错")
		return
	}
	res = strings.Replace(res, "\n", "", -1)
	res = strings.Replace(res, "\u0000", "", -1)
	res = strings.Split(res, "result: ")[1]
	middlewares.ResponseSuccess(ctx, "操作成功", res)
}

func (p *PlatonecliRequest) AddNode(ctx *webCtx.Context) {

	var NodeRequestInfo model.NodeOperateRequest

	if err := ctx.ShouldBind(&NodeRequestInfo); nil != err {
		middlewares.ResponseError(ctx, http.StatusBadRequest, err, "参数绑定失败！")
		return
	}

	if NodeRequestInfo.ExternalIP == "" || NodeRequestInfo.P2pPort == 0 || NodeRequestInfo.InternalIP == "" {
		err := exterror.ErrParameterInvalid
		middlewares.ResponseError(ctx, http.StatusBadRequest, err, "必填字段不能为空！")
		return

	}
	res, err := model.DefaultPlatonecliModel.AddNode(ctx, NodeRequestInfo)
	if nil != err {
		middlewares.ResponseError(ctx, http.StatusBadRequest, err, "操作失败！")
		return
	}

	middlewares.ResponseSuccess(ctx, "操作成功", res)
}

func (p *PlatonecliRequest) DeleteNode(ctx *webCtx.Context) {

	var deleteNodeRequestInfo model.DeleteNodeOperateRequest

	if err := ctx.ShouldBind(&deleteNodeRequestInfo); nil != err {
		middlewares.ResponseError(ctx, http.StatusBadRequest, err, "参数绑定失败！")
		return
	}

	if deleteNodeRequestInfo.Name == "" {
		err := exterror.ErrParameterInvalid
		middlewares.ResponseError(ctx, http.StatusBadRequest, err, "必填字段不能为空！")
		return

	}
	res, err := model.DefaultPlatonecliModel.DeleteNode(ctx, deleteNodeRequestInfo)
	if nil != err {
		middlewares.ResponseError(ctx, http.StatusBadRequest, err, "操作失败！")
		return
	}

	middlewares.ResponseSuccess(ctx, "操作成功", res)
}

func (p *PlatonecliRequest) DeployContract(ctx *webCtx.Context) {

	var DeployRequestInfo model.DeployRequest

	if err := ctx.ShouldBind(&DeployRequestInfo); nil != err {
		middlewares.ResponseError(ctx, http.StatusBadRequest, err, "参数绑定失败！")
		return
	}

	form, _ := ctx.MultipartForm()
	files := form.File["file"]

	for _, file := range files {
		log.Println(file.Filename)
		dst := path.Join("/tmp", file.Filename)
		ctx.SaveUploadedFile(file, dst)

		if strings.HasSuffix(file.Filename, "json") {
			DeployRequestInfo.AbiPath = dst
		} else if strings.HasSuffix(file.Filename, "wasm") || strings.HasSuffix(file.Filename, "evm") {
			DeployRequestInfo.CodePath = dst
		}

		if strings.HasSuffix(DeployRequestInfo.CodePath, "wasm") {
			DeployRequestInfo.Interpreter = "wasm"
		} else if strings.HasSuffix(DeployRequestInfo.CodePath, "evm") {
			DeployRequestInfo.Interpreter = "evm"
		}
	}

	if DeployRequestInfo.CodePath == "" || DeployRequestInfo.AbiPath == "" || DeployRequestInfo.Interpreter == "" {
		err := exterror.ErrParameterInvalid
		middlewares.ResponseError(ctx, http.StatusBadRequest, err, "必填字段不能为空！")
		return

	}
	res, err := model.DefaultPlatonecliModel.DeployContract(ctx, DeployRequestInfo)
	if nil != err {
		middlewares.ResponseError(ctx, http.StatusBadRequest, err, "操作失败！")
		return
	}

	middlewares.ResponseSuccess(ctx, "操作成功", res)
}

func (p *PlatonecliRequest) RoleSet(ctx *webCtx.Context) {

	var RoleSetInfo model.RoleSetRequest

	if err := ctx.ShouldBind(&RoleSetInfo); nil != err {
		middlewares.ResponseError(ctx, http.StatusBadRequest, err, "参数绑定失败！")
		return
	}

	if RoleSetInfo.Address == "" || RoleSetInfo.Role == "" {
		err := exterror.ErrParameterInvalid
		middlewares.ResponseError(ctx, http.StatusBadRequest, err, "必填字段不能为空！")
		return

	}
	res, err := model.DefaultPlatonecliModel.RoleSet(ctx, RoleSetInfo.Role, RoleSetInfo.Address)
	if nil != err {
		middlewares.ResponseError(ctx, http.StatusBadRequest, err, "操作失败！")
		return
	}

	middlewares.ResponseSuccess(ctx, "操作成功", res)
}

func runScript(s []string) (string, error) {
	out, err := exec.Command("/home/colin/Workspace/src/PlatONE-Go/release/linux/bin/ctool", s...).Output()
	return string(out), err
}
