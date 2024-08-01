package auth

import (
	"com.lh.basic/config"
	"com.lh.basic/crypto"
	serve "com.lh.service/locales"
	data "com.lh.service/pebble"
	"com.lh.service/tools"
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

func getTime(key string, value time.Duration) data.Config {
	data_dir := config.GetKey("DataDir")
	nowTime := time.Now()
	conf := data.Config{
		Path: fmt.Sprintf("%s/pebble/lh_%s", data_dir, key),
		Key:  key,
		DTSJ: nowTime.Format("20060102150405"),
		Data: data.DConf{
			"nowTime":        nowTime.Format("20060102150405"),
			"expirationTime": nowTime.Add(value).Format("20060102150405"),
		},
	}
	return conf
}

func CodeID(c *gin.Context) {
	id, err := crypto.UUID(c)
	if err != nil {
		tools.Code500(err.Error(), c)
	} else {
		conf := getTime("code", time.Minute*10)
		conf.Data["uuid"] = id
		res, err := data.AddLog(conf)

		if err != nil {
			tools.Code500(err.Error(), c)
		} else {
			tools.Code200(data.Conf{
				"code_id": res["id"],
				"id":      res["uuid"],
			}, c)
		}
	}
}

func GetUUID(c *gin.Context) {
	id, err := crypto.UUID(c)
	if err != nil {
		tools.Code500(err.Error(), c)
	} else {
		conf := getTime("uuid", time.Hour*6)
		conf.Data["uuid"] = id
		res, err := data.AddLog(conf)
		if err != nil {
			tools.Code500(err.Error(), c)
		} else {
			tools.Code200(data.Conf{
				"uuid_id": res["id"],
				"id":      res["uuid"],
			}, c)
		}
	}
}

func GetSID(c *gin.Context) {
	id := crypto.Md5("")
	conf := getTime("sid", time.Hour*6)
	conf.Data["key"] = id
	res, err := data.AddLog(conf)
	if err == nil {
		msg := serve.GetKey(c, []string{"errors", "query"})
		tools.Code500(msg, c)
		return
	}
	opts := crypto.AesConf{
		Key:   id,
		Value: res["id"].(string),
	}
	sid, err := crypto.AesEnCrypt(opts, c)
	if err != nil {
		tools.Code500(err.Error(), c)
		return
	}
	tools.Code200(data.DConf{
		"sid":     sid,
		"code_id": res["id"].(string),
	}, c)
	return
}
