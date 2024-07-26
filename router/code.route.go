package router

import (
	"com.lh.auth/code"
	"github.com/gin-gonic/gin"
)

func CodeRoute(app *gin.RouterGroup) *gin.RouterGroup {
	app.GET("/code", func(context *gin.Context) {
		go code.GetCode(context)
	})
	app.GET("/del", func(context *gin.Context) {
		go code.Del(context)
	})
	return app
}
