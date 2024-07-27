package main

import (
	"com.lh.auth/locales"
	"com.lh.auth/router"
	"com.lh.basic/config"
	"com.lh.service/tools"
	"com.lh.service/yugabyte"
	"fmt"
	"github.com/gin-gonic/gin"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU() * 2)
	app := gin.Default()
	config.InitConfig("com.lh.auth")
	configs := config.GetConfig("auth")
	yugabyte.InitConfig()
	app.Use(tools.Cors())
	//app.Use(tools.MiddleWare(configs))
	locales.Init()
	router.Router(app)
	app.Run(fmt.Sprintf("%s:%s", configs.Host, configs.Port))
}
