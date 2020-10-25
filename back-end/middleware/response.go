// pkg/utils/middleware/responseStatus.go
package middlewares

import (
	webCtx "data-manager/web/context"
	"encoding/json"
)

/**
封装服务器返回状态吗
*/

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

//错误状态处理
func ResponseError(ctx *webCtx.Context, code int, err error, msg string) {
	resp := &Response{Code: code, Msg: msg, Data: err}
	ctx.JSON(code, resp)
	response, _ := json.Marshal(resp)
	ctx.Set("response", string(response))
	ctx.AbortWithError(code, err)
}

//正确状态处理
func ResponseSuccess(ctx *webCtx.Context, msg string, data interface{}) {
	resp := &Response{Code: 200, Msg: msg, Data: data}
	ctx.JSON(200, resp)
	response, _ := json.Marshal(resp)
	ctx.Set("response", string(response))
}
