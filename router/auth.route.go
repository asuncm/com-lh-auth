package router

import (
	"com.lh.auth/src/auth"
	"github.com/gin-gonic/gin"
)

func AuthRoute(app *gin.RouterGroup) *gin.RouterGroup {
	app.GET("/code/clientId", func(c *gin.Context) {
		auth.CodeID(c)
	})
	return app
}
