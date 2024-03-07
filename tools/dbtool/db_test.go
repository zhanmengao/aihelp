package dbtool

import (
	"path"
	"testing"

	"forevernine.com/midplat/base_server/tools/f9/dbtool/internal/gen"
)

func TestDB(t *testing.T) {
	gen.Run("D:/go_workspace/MidPlatProto/go/pb", "forevernine.com/midplat/midplat_proto/go/pb", "database", "midplat", "")
}

func TestAAA(t *testing.T) {
	t.Log(path.Dir("forevernine.com/midplat/base_server/tools/f9/dbtool/internal/gen"))
}
