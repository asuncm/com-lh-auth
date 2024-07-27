package auth

import (
	"com.lh.basic/config"
	"com.lh.basic/crypto"
	data "com.lh.service/pebble"
	"com.lh.service/tools"
	"com.lh.service/yugabyte"
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

func CodeID(c *gin.Context) {
	data_dir := config.Config.DataDir
	nowTime := time.Now()
	id, err := crypto.UUID(c)
	if err != nil {
		tools.Code500(err.Error(), c)
	} else {
		conf := data.Config{
			Path: fmt.Sprintf("%s/pebble/auth", data_dir),
			Key:  "uuid",
			DTSJ: nowTime.UnixNano(),
			Data: data.DConf{
				"uuid":           id,
				"nowTime":        nowTime.UnixNano(),
				"expirationTime": nowTime.Add(time.Minute * 10).UnixNano(),
			},
		}
		//res, err := data.Append(config)
		fmt.Println(data_dir, nowTime, id, conf)
		yugabyte.Ping(yugabyte.Config{
			Name: "yugabyte",
			DB:   "allkic",
		})
		if err != nil {
			tools.Code500(err.Error(), c)
		} else {
			tools.Code200("res.Data", c)
		}
	}
}
