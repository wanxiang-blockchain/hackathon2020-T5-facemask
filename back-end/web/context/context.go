package context

import (
	dbCtx "data-manager/db/context"
	"github.com/gin-gonic/gin"
)

type Context struct {
	*gin.Context
	DBCtx *dbCtx.Context
}

func New() *Context {
	return &Context{}
}

func (this *Context) SetContext(ctx *gin.Context) {
	this.Context = ctx
}

func (this *Context) SetDBContext(ctx *dbCtx.Context) {
	this.DBCtx = ctx
}

//func (this *Context) JSON(ctx *dbCtx.Context, code int, obj interface{}) {
//	var msg interface{}
//	if err, ok := obj.(error); ok {
//		var extErr *exterror.ExtError
//		if errors.As(err, &extErr) {
//			msg = NewFailedResponse(extErr)
//			middlewares.ResponseError(ctx, http.StatusBadRequest, err, "errors")
//		} else {
//			msg = NewFailedResponse(exterror.NewError(exterror.ErrCodeInterval, err.Error()))
//		}
//	} else {
//		msg = NewOkResponse(obj)
//	}
//
//	this.Context.JSON(code, msg)
//}
