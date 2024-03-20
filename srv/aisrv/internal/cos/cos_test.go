package cos

import (
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/zhanmengao/aihelp/config"
	"github.com/zhanmengao/aihelp/global"
	"github.com/zhanmengao/aihelp/manifest/protobuf/gendb"
	"github.com/zhanmengao/aihelp/manifest/protobuf/pb"
	"testing"
)

func TestMain(m *testing.M) {
	global.PikaDB = gendb.NewPikaDBCache()
	config.GptConfig.V3ApiKey = "sk-WYd9zVPgx0ywV7CW634906Dc3b204aD3Bf6070E3Ef3426Cc"
	config.GptConfig.V3Host = "https://one.aiskt.com"
	m.Run()
}

func TestCat(t *testing.T) {
	srv := &cosService{}
	rsp, err := srv.CosPlay(gctx.New(), &pb.CosPlayReq{
		RoleType: pb.CosPlayCat,
		RoleDesc: "一只叫肥特的小猫",
	})
	if err != nil {
		t.Error(err)
		t.Failed()
	} else {
		for _, msg := range rsp.Message {
			t.Log(msg.Content)
		}

	}
}

func TestGirl(t *testing.T) {
	srv := &cosService{}
	rsp, err := srv.CosPlay(gctx.New(), &pb.CosPlayReq{
		RoleType: pb.CosPlayGirlFriend,
		RoleDesc: "原神的宵宫",
	})
	if err != nil {
		t.Error(err)
		t.Failed()
	} else {
		for _, msg := range rsp.Message {
			t.Log(msg.Content)
		}

	}
}
