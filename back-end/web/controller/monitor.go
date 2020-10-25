package controller

import (
	"data-manager/exterror"
	middlewares "data-manager/middleware"
	"data-manager/model"
	"data-manager/model/monitor"
	webCtx "data-manager/web/context"
	"net/http"
)

type monitorConctroller struct{}

var DefaultMonitorController = &monitorConctroller{}

func (this *monitorConctroller) NewNode(ctx *webCtx.Context) {
	var nodeinfo model.Node
	if err := ctx.ShouldBind(&nodeinfo); nil != err {
		middlewares.ResponseError(ctx, http.StatusBadRequest, err, "参数绑定失败！")
		return
	}
	if nodeinfo.MonitorIp == "" || nodeinfo.MonitorPort == "" {
		err := exterror.ErrParameterInvalid
		middlewares.ResponseError(ctx, http.StatusBadRequest, err, "ip和端口号不能为空！")
		return

	}
	res, err := monitor.DefaultMonitorModel.New(ctx, nodeinfo)
	if nil != err {
		middlewares.ResponseError(ctx, http.StatusBadRequest, err, "创建节点失败！")
		return
	}

	middlewares.ResponseSuccess(ctx, "创建节点成功！", res)
	return

}

func (this *monitorConctroller) StartNode(ctx *webCtx.Context) {
	var nodeInfo model.NodeOp
	if err := ctx.ShouldBind(&nodeInfo); nil != err {
		middlewares.ResponseError(ctx, http.StatusBadRequest, err, "参数绑定失败！")
		return
	}
	if nodeInfo.PublicKey == "" || nodeInfo.Groupid == "" {
		err := exterror.ErrParameterInvalid
		middlewares.ResponseError(ctx, http.StatusBadRequest, err, "公钥和groupid不能为空！")
		return

	}
	err := monitor.DefaultMonitorModel.StartNode(ctx, nodeInfo)
	if nil != err {
		middlewares.ResponseError(ctx, http.StatusBadRequest, err, "启动节点失败！")
		return
	}

	middlewares.ResponseSuccess(ctx, "启动节点成功！", nil)
	return
}

func (this *monitorConctroller) StopNode(ctx *webCtx.Context) {
	var nodeInfo model.NodeOp
	if err := ctx.ShouldBind(&nodeInfo); nil != err {
		middlewares.ResponseError(ctx, http.StatusBadRequest, err, "参数绑定失败！")
		return
	}
	if nodeInfo.PublicKey == "" || nodeInfo.Groupid == "" {
		err := exterror.ErrParameterInvalid
		middlewares.ResponseError(ctx, http.StatusBadRequest, err, "公钥和groupid不能为空！")
		return

	}
	err := monitor.DefaultMonitorModel.StopNode(ctx, nodeInfo)
	if nil != err {
		middlewares.ResponseError(ctx, http.StatusBadRequest, err, "关闭节点失败！")
		return
	}

	middlewares.ResponseSuccess(ctx, "关闭节点成功！", nil)
	return
}

func (this *monitorConctroller) RestartNode(ctx *webCtx.Context) {
	var nodeInfo model.NodeOp
	if err := ctx.ShouldBind(&nodeInfo); nil != err {
		middlewares.ResponseError(ctx, http.StatusBadRequest, err, "参数绑定失败！")
		return
	}
	if nodeInfo.PublicKey == "" || nodeInfo.Groupid == "" {
		err := exterror.ErrParameterInvalid
		middlewares.ResponseError(ctx, http.StatusBadRequest, err, "公钥和groupid不能为空！")
		return

	}
	err := monitor.DefaultMonitorModel.RestartNode(ctx, nodeInfo)
	if nil != err {
		middlewares.ResponseError(ctx, http.StatusBadRequest, err, "关闭节点失败！")
		return
	}

	middlewares.ResponseSuccess(ctx, "关闭节点成功！", nil)
	return
}
