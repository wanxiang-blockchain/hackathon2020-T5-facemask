package middlewares

import (
	"github.com/gin-gonic/gin"
)

func SuperAdminAuthMiddleWare(c *gin.Context) {
	// tmpSD, _ := c.Get(gin_session.SessionContextName)
	// sd := tmpSD.(gin_session.SessionData)
	// value, err := sd.Get("isSuperAdminLogin")
	// if err != nil {
	// 	resp := &Response{Code: 400, Msg: "NOPERMISSION", Data: nil}

	// 	c.JSON(400, resp)
	// 	c.AbortWithStatus(400)
	// 	return
	// }
	// fmt.Println(value)
	// isLogin, ok := value.(bool)
	// if !ok {
	// 	resp := &Response{Code: 400, Msg: "NOPERMISSION", Data: nil}
	// 	c.JSON(400, resp)
	// 	c.AbortWithStatus(400)
	// 	return
	// }
	// fmt.Println(isLogin)
	// if !isLogin {
	// 	resp := &Response{Code: 400, Msg: "NOPERMISSION", Data: nil}
	// 	c.JSON(400, resp)
	// 	c.AbortWithStatus(400)
	// 	return
	// }
	c.Next()
}
