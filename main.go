package main

import (
	auth "com.lh.auth/locales"
	"com.lh.auth/router"
	"com.lh.basic/config"
	basic "com.lh.basic/locales"
	serve "com.lh.service/locales"
	"com.lh.service/tools"
	"com.lh.service/yugabyte"
	user "com.lh.user/locales"
	web "com.lh.web.service/locales"
	"fmt"
	"github.com/gin-gonic/gin"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU() * 2)
	app := gin.Default()
	app.Use(gin.Logger())
	app.Use(gin.Recovery())
	config.InitConfig("com.lh.auth")
	root := config.GetKey("Root").(string)
	configs := config.GetServe("auth")
	yugabyte.InitConfig()
	app.Use(tools.Cors())
	//app.Use(tools.MiddleWare(configs))
	auth.Init(root)
	basic.Init(root)
	user.Init(root)
	web.Init(root)
	serve.Init(root)
	router.Router(app)
	app.Run(fmt.Sprintf("%s:%s", configs.Host, configs.Port))
}
