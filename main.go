package main

import (
	"com.lh.auth/locales"
	router2 "com.lh.auth/router"
	"com.lh.service/src/tools"
	"fmt"
	"github.com/gin-gonic/gin"
	"strings"
)

func Config() (tools.MiddleConf, error) {
	platform := tools.Platform("")
	pathname := tools.GetPath("LHPATH", fmt.Sprintf("%s%s%s", "config/", platform.Env, ".config.yaml"))
	configs, err := tools.Yaml(pathname)
	if err != nil {
		return tools.MiddleConf{}, err
	}
	devServe := configs.Services["auth"]
	root := configs.Root
	database := fmt.Sprintf("%s%s", configs.Database, "/pebble")
	return tools.MiddleConf{
		Platform:  platform.Platform,
		Serve:     "auth",
		Root:      root,
		Host:      devServe.Host,
		Port:      devServe.Port,
		DataCache: database,
		DataPort:  devServe.DataPort,
	}, err
}

func main() {
	router := gin.Default()
	configs, _ := Config()
	router.Use(tools.MiddleWare(configs))
	locales.Init()
	router2.Router(router)
	address := []string{configs.Host, configs.Port}
	router.Run(strings.Join(address, ":"))
}
