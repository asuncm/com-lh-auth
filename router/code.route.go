package router

import "github.com/gin-gonic/gin"

func CodeRoute(app *gin.RouterGroup) *gin.RouterGroup {
	app.GET("/code", func(context *gin.Context) {

	})
	return app
}
