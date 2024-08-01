package code

import (
	"com.lh.service/tools"
	"com.lh.service/yugabyte"
	"fmt"
	"github.com/gin-gonic/gin"
)

func GetCode(c *gin.Context) {
	//data_dir, _ := c.Get("DataDir")
	//config := data.Config{
	//	Path: data_dir.(string),
	//	Option: data.Option{
	//		ID:        "1718871283211903300",
	//		Min:       1,
	//		Key:       "uuid_log",
	//		Prefix:    "uuid",
	//		StartTime: "1718815149766443300",
	//	},
	//}
	//data, err := data.GetKey(config)
	//fmt.Println(data, err, "--------------")
	//var buf auth.Code
	//if err == nil {
	//	value := gob.NewDecoder(bytes.NewReader(data.([]byte)))
	//	err = value.Decode(&buf)
	//}
	//buf.Pd = crypto.Md5("Allkic20230619@@")

	type cd struct {
		id   int16
		name string
	}
	//pool, err := yugabyte.PoolDB("Allkic", "allkic")
	//conn, err := yugabyte.OpenDB("Yugabyte", "yugabyte")
	//fmt.Println(conn, err)
	//opts := yugabyte.Config{
	//	Name: "Allkic",
	//	DB:   "allkic",
	//	Key:  "allkic",
	//}
	p := yugabyte.Ping(yugabyte.Config{
		Name: "yugabyte",
		DB:   "yugabyte",
	})
	fmt.Println(p, "lllpppp====")
	defer tools.Code200("==========", c)
}

func Del(c *gin.Context) {
	//data_dir := config2.Config.DataDir
	//nowTime := time.Now()
	//id, err := crypto.UUID(c)
	//if err != nil {
	//
	//}
	//config := data.Config{
	//	Path: fmt.Sprintf("%s/pebble/auth", data_dir),
	//	Key:  "uuid",
	//	DTSJ: nowTime.UnixNano(),
	//	Data: data.DConf{
	//		"uuid":           id,
	//		"nowTime":        nowTime.UnixNano(),
	//		"expirationTime": nowTime.Add(time.Minute * 10).UnixNano(),
	//	},
	//}
	tools.Code200(map[string]interface{}{
		"isBool": "p",
		"sss":    "sssss",
	}, c)
}
