package framework

import "forevernine.com/midplat/base_libs/idgen"

const (
	KeyNameUID = "base_uid"
)

var (
	uidGen *idgen.TIDGen
)

func init() {
	uidGen, _ = idgen.NewIDGen()
}
