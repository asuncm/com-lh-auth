package code

import (
	"bytes"
	"com.lh.auth/src/auth"
	"com.lh.basic/src/crypto"
	data "com.lh.service/src/pebble"
	"com.lh.service/src/tools"
	"com.lh.service/src/yugabyte"
	"encoding/gob"
	"fmt"
	"github.com/gin-gonic/gin"
)

func Code(c *gin.Context) {
	paths := tools.Pathname(c, "/")
	fmt.Println("data", "err", "--------------")
	config := data.Config{
		Path: paths.Path,
		Option: data.Option{
			ID:        "1718871283211903300",
			Min:       1,
			Key:       "uuid_log",
			Prefix:    "uuid",
			StartTime: "1718815149766443300",
		},
	}
	data, err := data.GetKey(config)
	fmt.Println(data, err, "--------------")
	var buf auth.Code
	if err == nil {
		value := gob.NewDecoder(bytes.NewReader(data.([]byte)))
		err = value.Decode(&buf)
	}
	buf.Pd = crypto.Md5("Allkic20230619@@")

	//err = yugabyte.TABLE(`
	//	CREATE TABLE users (
	//	    id serial PRIMARY KEY,
	//	    name text NOT NULL
	//	)
	//`)
	//fmt.Println(err, "创建数据库")

	type cd struct {
		id   int16
		name string
	}
	//dm, err := yugabyte.ADD(`
	//	INSERT INTO users (id, name) VALUES (1, 'ceshi')`)
	//dm, err := yugabyte.ADD("Allkic", "allkic", []string{"id", "name"}, [][]any{{11, "ddd"}})
	dm, err := yugabyte.OpenDB("Allkic", "allkic")
	fmt.Println(dm, err, "errrrrrrrrrrrrrrr=============")
	tools.Code200(buf, c)
}

func Del(c *gin.Context) {
	paths := tools.Pathname(c, "/")
	fmt.Println("data", "err", "--------------")
	config := data.Config{
		Path: paths.Path,
		Option: data.Option{
			ID:        "1718871283211903300",
			Min:       1,
			Key:       "uuid_log",
			Prefix:    "uuid",
			StartTime: "1718815149766443300",
		},
	}
	db, _ := data.OpenDB(config.Path, config.Config)
	err := data.Delkey(db, "uuid_"+"1718871283211903300")
	db.Close()
	fmt.Println(err)
	tools.Code200(nil, c)
}
