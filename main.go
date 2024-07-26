package main

import (
	"com.lh.auth/locales"
	"com.lh.auth/router"
	"com.lh.basic/config"
	"com.lh.service/pgx"
	"com.lh.service/tools"
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	app := gin.Default()
	config.InitConfig("com.lh.auth")
	configs := config.GetConfig("auth")
	pgx.InitConfig()
	app.Use(tools.Cors())
	//app.Use(tools.MiddleWare(configs))
	locales.Init()
	router.Router(app)
	app.Run(fmt.Sprintf("%s:%s", configs.Host, configs.Port))
}
