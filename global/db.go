package global

import "github.com/zhanmengao/aihelp/manifest/protobuf/gendb"

var (
	PikaDB gendb.IPikaDB
)

func InitDB() (err error) {
	PikaDB = gendb.NewPikaDBRedis()
	return
}
