package router

import (
	"com.lh.auth/auth"
	"github.com/gin-gonic/gin"
)

func AuthRoute(app *gin.RouterGroup) *gin.RouterGroup {
	app.GET("/code/code", func(c *gin.Context) {
		auth.CodeID(c)
		return
	})
	app.GET("/sid", func(c *gin.Context) {
		auth.GetSID(c)
		return
	})
	app.GET("/uuid", func(c *gin.Context) {
		auth.GetUUID(c)
		return
	})
	return app
}
