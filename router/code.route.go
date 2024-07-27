package router

import (
	"com.lh.auth/code"
	"github.com/gin-gonic/gin"
)

func CodeRoute(app *gin.RouterGroup) *gin.RouterGroup {
	go app.GET("/code", func(context *gin.Context) {
		code.GetCode(context)
	})
	go app.GET("/del", func(context *gin.Context) {
		code.Del(context)
	})
	return app
}
