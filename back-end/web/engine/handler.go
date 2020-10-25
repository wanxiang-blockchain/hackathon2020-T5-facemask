package engine

import (
	"data-manager/db"
	dbContext "data-manager/db/context"
	webContext "data-manager/web/context"
	"github.com/gin-gonic/gin"
)

type HandlerFunc func(ctx *webContext.Context)

func NewHandler(webC *webContext.Context, handlers ...HandlerFunc) func(*gin.Context) {
	//fmt.Printf("new handler out")
	//fmt.Printf("%#v\n", webC)
	return func(ctx *gin.Context) {
		//fmt.Printf("new handler")
		//
		//fmt.Printf("%#v\n", webC)

		webCtx := webC

		webCtx.SetContext(ctx)

		dbCtx := dbContext.New(db.DefaultDB)
		webCtx.SetDBContext(dbCtx)

		for _, h := range handlers {
			h(webCtx)
		}
	}
}
