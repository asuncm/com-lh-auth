package router

import (
	"com.lh.auth/crypto"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AuthRoute(app *gin.RouterGroup) *gin.RouterGroup {
	app.GET("/code/test", func(c *gin.Context) {
		text, _ := crypto.AesEnCrypt(crypto.AesConf{
			Key:   "1234567890qwerty",
			Value: "456qwertyuiop",
		}, "")
		dd, _ := crypto.AesDeCrypt(crypto.AesConf{
			Key:   "1234567890qwerty",
			Value: "cc68a0f022ba23ea3c3894636ac8829d",
		}, "")
		fmt.Println(text, "text", dd)
		c.String(http.StatusOK, "==============>>>>", text, "======000", dd)
	})
	return app
}
