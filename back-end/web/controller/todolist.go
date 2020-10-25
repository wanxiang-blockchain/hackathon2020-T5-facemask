package controller

import (
	middlewares "data-manager/middleware"
	"data-manager/model"
	webCtx "data-manager/web/context"
	"net/http"
)

//func init() {
//	webEngine.Default.GET("/header", webEngine.NewHandler(DefaultBlockController.AddHeader))
//}
//todo 放到router中去做

type blockController struct{}

var DefaultBlockController = &blockController{}

func (this *blockController) DoSomething(ctx *webCtx.Context) {
	//var h Header
	//err := ctx.ShouldBindJSON(&h)
	//if nil != err {
	//	ctx.JSON(http.StatusBadRequest, err)
	//	return
	//}
	//
	err := model.DefaultUserModel.Test(ctx)
	if nil != err {
		middlewares.ResponseError(ctx, http.StatusBadRequest, err, "errors")
		ctx.JSON(http.StatusBadRequest, err)
	}

	middlewares.ResponseSuccess(ctx, "请求成功", nil)

}
