package middlewares

import (
	gin_session "data-manager/web/session"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware(c *gin.Context) {
	tmpSD, _ := c.Get(gin_session.SessionContextName)
	sd := tmpSD.(gin_session.SessionData)
	value, err := sd.Get("isLogin")
	if err != nil {
		resp := &Response{Code: 400, Msg: "NOPERMISSION", Data: nil}

		c.JSON(400, resp)
		c.AbortWithStatus(400)
		return
	}
	isLogin, ok := value.(bool)
	if !ok {
		resp := &Response{Code: 400, Msg: "NOPERMISSION", Data: nil}

		c.JSON(400, resp)
		c.AbortWithStatus(400)
		return
	}
	if !isLogin {
		resp := &Response{Code: 400, Msg: "NOPERMISSION", Data: nil}

		c.JSON(400, resp)
		c.AbortWithStatus(400)
		return
	}
	c.Next()
}
