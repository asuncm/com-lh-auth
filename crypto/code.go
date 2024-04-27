package crypto

import "com.lh.service/badger"

func ListId(value int64, pathname string) (badger.DataConf, error) {
	node, _ := NewWorker(value, pathname)
	id := node.GetId()
	key, _ := UUID(pathname)
	md5 := Md5(key)
	return badger.DataConf{
		id:    id,
		key:   key,
		"mId": md5,
	}, nil
}
