package config

import (
	"com.lh.auth/locales"
	"com.lh.basic/config"
	basic "com.lh.basic/locales"
	serve "com.lh.service/locales"
	user "com.lh.user/locales"
	web "com.lh.web.service/locales"
)

func Init() {
	config.InitConfig("com.lh.auth")
	root := config.GetKey("Root").(string)
	locales.Init(root)
	basic.Init(root)
	serve.Init(root)
	user.Init(root)
	web.Init(root)
}
