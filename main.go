package main

import (
	"com.lh.auth/locales"
	router2 "com.lh.auth/router"
	"com.lh.service/tools"
	"fmt"
	"github.com/gin-gonic/gin"
	"strings"
)

func Config() (tools.MiddleConf, error) {
	platform := tools.Platform("")
	root := tools.GetPath("LHPATH", "")
	pathname := fmt.Sprintf("%s%s%s%s", root, "/config/", platform.Env, ".config.yaml")
	configs, err := tools.Yaml(pathname)
	if err != nil {
		return tools.MiddleConf{}, err
	}
	devServe := configs.Services["auth"]
	database := tools.GetPath(configs.Database, "pebble/auth")
	return tools.MiddleConf{
		Platform: platform.Platform,
		Serve:    fmt.Sprintf("%s%s", root, "/com.lh.auth"),
		Root:     root,
		Host:     devServe.Host,
		Port:     devServe.Port,
		DataDir:  database,
	}, err
}

func main() {
	router := gin.Default()
	configs, _ := Config()
	router.Use(tools.Cors())
	router.Use(tools.MiddleWare(configs))
	locales.Init()
	router2.Router(router)
	address := []string{configs.Host, configs.Port}
	router.Run(strings.Join(address, ":"))
}
