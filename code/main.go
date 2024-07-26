package code

import (
	data "com.lh.service/pebble"
	"com.lh.service/tools"
	"fmt"
	"github.com/gin-gonic/gin"
)

func GetCode(c *gin.Context) {
	//data_dir, _ := c.Get("DataDir")
	fmt.Println("data", "err", "--------------")
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
	//pool, err := pgx.PoolDB("Allkic", "allkic")
	//conn, err := pgx.OpenDB("Yugabyte", "pgx")
	//fmt.Println(conn, err)
	//opts := pgx.Config{
	//	Name: "Allkic",
	//	DB:   "allkic",
	//	Key:  "allkic",
	//}
	fmt.Println("errrrrrrrrrrrrrrr=============", "boo")
	tools.Code200(nil, c)
}

func Del(c *gin.Context) {
	data_dir, _ := c.Get("DataDir")
	fmt.Println("data", "err", "--------------")
	config := data.Config{
		Path: data_dir.(string),
	}
	//db, _ := data.OpenDB(config.Path, config.Config)
	////err := data.Delkey(db, "uuid_"+"1718871283211903300")
	//db.Close()
	fmt.Println("err", config)
	tools.Code200(nil, c)
}
