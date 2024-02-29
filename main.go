package main

import (
	_ "github.com/zhanmengao/aihelp/internal/packed"

	"github.com/gogf/gf/v2/os/gctx"

	"github.com/zhanmengao/aihelp/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}
