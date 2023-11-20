package main

import (
	router2 "com.lh.auth/src/router"
	"com.lh.service/src/config/middleware"
	"com.lh.service/src/config/yaml"
	"com.lh.service/src/tools"
	"github.com/gin-gonic/gin"
	"os"
	"strings"
)

func Config() (middleware.MiddleConf, error) {
	platform := tools.Platform("")
	dir, _ := os.Getwd()
	dir = strings.Replace(dir, "\\", "/", -1)
	paths := []string{dir, "config", platform.Env + ".config.yaml"}
	pathname := strings.Join(paths, "/")
	configs, err := yaml.Yaml(pathname)
	if err != nil {
		return middleware.MiddleConf{}, err
	}
	devServe := configs.Services["auth"]
	root := configs.Root[platform.Platform]
	database := configs.Database["badger"]
	return middleware.MiddleConf{
		Platform:  platform.Platform,
		Serve:     "auth",
		Root:      root,
		Host:      devServe.Host,
		Port:      devServe.Port,
		DataCache: database[platform.Platform],
		DataPort:  devServe.DataPort,
	}, err
}

func main() {
	router := gin.Default()
	configs, _ := Config()
	router.Use(middleware.MiddleWare(configs))
	router2.Router(router)
	address := []string{configs.Host, configs.Port}
	router.Run(strings.Join(address, ":"))
}
