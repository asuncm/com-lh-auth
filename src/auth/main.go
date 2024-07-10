package auth

import (
	"com.lh.basic/src/crypto"
	data "com.lh.service/src/pebble"
	"com.lh.service/src/tools"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
)

type Code struct {
	ID             string `json:"id"`
	NowTime        int64  `json:"nowTime"`
	ExpirationTime int64  `json:"expirationTime"`
	Pd             string `json:"pd"`
}

func CodeID(c *gin.Context) {
	paths := tools.Pathname(c, "/")
	nowTime := time.Now()
	id, err := crypto.UUID(c)
	if err != nil {
		tools.Code500(err.Error(), c)
	} else {
		lists := Code{
			ID:             id,
			NowTime:        nowTime.UnixNano(),
			ExpirationTime: nowTime.Add(time.Minute * 10).UnixNano(),
		}
		config := data.Config{
			Path: paths.Path,
			Option: data.Option{
				ID:        strconv.FormatInt(lists.NowTime, 10),
				Min:       10,
				Key:       "uuid_log",
				Prefix:    "uuid",
				StartTime: nowTime.Format("200601021504"),
			},
		}
		err = data.Append(config, lists)
		fmt.Println(config, err, "PPPPPP====")
		tools.Code200(lists, c)
	}
}
