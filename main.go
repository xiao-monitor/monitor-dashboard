package main

import (
	_ "monitor-dashboard/internal/packed"

	"github.com/gogf/gf/v2/os/gctx"

	"monitor-dashboard/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}
