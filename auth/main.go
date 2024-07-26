package auth

import (
	"com.lh.basic/crypto"
	data "com.lh.service/pebble"
	"com.lh.service/pgx"
	"com.lh.service/tools"
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

func CodeID(c *gin.Context) {
	data_dir, _ := c.Get("DataDir")
	nowTime := time.Now()
	id, err := crypto.UUID(c)
	if err != nil {
		tools.Code500(err.Error(), c)
	} else {
		config := data.Config{
			Path:     data_dir.(string),
			Key:      "uuid",
			Max:      100,
			Duration: time.Minute * 10,
			Data: data.DConf{
				"uuid":           id,
				"nowTime":        nowTime.UnixNano(),
				"expirationTime": nowTime.Add(time.Minute * 10).UnixNano(),
			},
		}
		//res, err := data.Append(config)
		fmt.Println(config)
		pgx.Ping(pgx.Config{
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
