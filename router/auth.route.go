package router

import (
	"com.lh.auth/crypto"
	"com.lh.service/badger"
	"com.lh.service/tools"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AuthRoute(app *gin.RouterGroup) *gin.RouterGroup {
	app.GET("/code/clientId", func(c *gin.Context) {
		workId, _ := c.Get("WorkerID")
		node, _ := crypto.NewWorker(int64(any(workId).(int8)), "")
		id := node.GetId()
		pathname := tools.Pathname(c, "/")
		badger.SetWithTTL(pathname, id, badger.DataConf{}, 30*60)
		fmt.Println(30>>4, 30<<6)
		c.JSON(http.StatusOK, tools.Code200(nil, id))
	})
	return app
}
