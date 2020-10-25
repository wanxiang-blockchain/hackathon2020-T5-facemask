package utils

import (
	middlewares "data-manager/middleware"
	webCtx "data-manager/web/context"
	"fmt"
	"net"
	"net/http"
	"os"
)

func GetIp() string{
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	for _, addr := range addrs {
		if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				fmt.Println(ipnet.IP.String())
				return ipnet.IP.String()
			}
		}
	}
	return ""
}
func ParamBind(ctx *webCtx.Context, param interface{}) {
	if err := ctx.ShouldBind(&param); nil != err {
		middlewares.ResponseError(ctx, http.StatusBadRequest, err, "参数绑定失败! ")
		return
	}
}

