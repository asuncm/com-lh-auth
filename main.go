package main

import (
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
	app.Use(gin.Logger())
	app.Use(gin.Recovery())
	opts := config.InitConfig("com.lh.auth")

	yugabyte.InitConfig()
	app.Use(tools.Cors())
	//app.Use(tools.MiddleWare(configs))
	fmt.Println(opts, "oopp---------------s")
	app.Run(fmt.Sprintf("%s:%s", opts["host"].(string), opts["port"].(string)))
}
