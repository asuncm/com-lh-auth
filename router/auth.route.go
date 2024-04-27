package router

import (
	"com.lh.auth/crypto"
	"com.lh.service/badger"
	"com.lh.service/tools"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AuthRoute(app *gin.RouterGroup) *gin.RouterGroup {
	app.GET("/code/clientId", func(c *gin.Context) {
		workId, _ := c.Get("WorkerID")
		list, _ := crypto.ListId(int64(any(workId).(int8)), "")
		pathname := tools.Pathname(c, "/")
		badger.SetWithTTL(pathname, any(list["id"]).(string), list, 30*60)
		c.JSON(http.StatusOK, tools.Code200(nil, list["id"]))
	})
	return app
}
