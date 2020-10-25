package engine

import (
	"data-manager/config"
	"github.com/gin-gonic/gin"
)

type Engine struct {
	*gin.Engine
}

var Default *Engine

func init() {
	gin.SetMode(config.Config.HttpConf.Mode)

	Default = New()

	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	Default.Use(gin.Recovery())
}

func New() *Engine {
	return &Engine{Engine: gin.New()}
}

// Note: this method will block the calling goroutine indefinitely unless an error happens.
func (this *Engine) Run(addr string) {
	this.Engine.Run(addr)

	panic("failed to web engine start.")
}
